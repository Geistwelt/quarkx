/**
* Author: Xiangyu Wu
* Date: 2023-06-23
* From: hyperledger/fabric/bccsp/utils/ecdsa.go
 */

package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"
)

// ECDSASignature 椭圆曲线签名，生成签名的过程如下所示：
//  1. 随机选择一个整数 k，满足 1 < k < p-1;
//  2. 计算点 R = kG;
//  3. 计算 r = x(R), 其中 x(R) 表示点 R 的 x 坐标;
//  4. 计算 s = (H(m) + d·r) · k^-1 mod (p-1).
type ECDSASignature struct {
	R, S *big.Int
}

// 当使用 base point 阶的一半时，可以增加签名方案的安全性，因为这样可以增加攻击者的计算次数，
// 理由是当阶减小一半时，就可能存在多个解，这会增加攻击者寻找合适解的难度。
var (
	// 将 base point 的阶右移一位
	curveHalfOrders = map[elliptic.Curve]*big.Int{
		// 13479973333575319897333507543509812970228903857212195860841361184030
		elliptic.P224(): new(big.Int).Rsh(elliptic.P224().Params().N, 1),
		// 57896044605178124381348723474703786764998477612067880171211129530534256022184
		elliptic.P256(): new(big.Int).Rsh(elliptic.P256().Params().N, 1),
		// 19701003098197239606139520050071806902539869635232723333973452639813829699556631784699478154076147456777216826971321
		elliptic.P384(): new(big.Int).Rsh(elliptic.P384().Params().N, 1),
		// 3432398830065304857490950399540696608634717650071652704697231729592771591698827697122528873166608598766481998185681660556932384306220190170186404446353502724
		elliptic.P521(): new(big.Int).Rsh(elliptic.P521().Params().N, 1),
	}
)

// GetCurveHalfOrdersAt 根据给定的椭圆曲线，返回椭圆曲线的 base point 的阶的一半。
func GetCurveHalfOrdersAt(c elliptic.Curve) *big.Int {
	return new(big.Int).Set(curveHalfOrders[c])
}

// MarshalECDSASignature 将椭圆曲线签名转化为用于描述数据结构的标记语言。
func MarshalECDSASignature(r, s *big.Int) ([]byte, error) {
	return asn1.Marshal(ECDSASignature{R: r, S: s})
}

// UnmarshalECDSASignature 反序列化椭圆曲线签名，先后返回签名的 R 和 S。
func UnmarshalECDSASignature(raw []byte) (*big.Int, *big.Int, error) {
	sig := new(ECDSASignature)
	if _, err := asn1.Unmarshal(raw, sig); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal signature [%v]", err)
	}

	if sig.R == nil {
		return nil, nil, errors.New("invalid signature, R must be different from nil")
	}

	if sig.S == nil {
		return nil, nil, errors.New("invalid signature, S must be different from nil")
	}

	if sig.R.Sign() != 1 {
		return nil, nil, errors.New("invalid signature, R must be larger than 0")
	}

	if sig.S.Sign() != 1 {
		return nil, nil, errors.New("invalid signature, S must be larger than 0")
	}

	return sig.R, sig.S, nil
}

// SignatureToLowS 让椭圆曲线签名的 S 部分小于椭圆曲线的 base point 的阶的一半。
func SignatureToLowS(k *ecdsa.PublicKey, signature []byte) ([]byte, error) {
	r, s, err := UnmarshalECDSASignature(signature)
	if err != nil {
		return nil, err
	}

	s, err = ToLowS(k, s)
	if err != nil {
		return nil, err
	}

	return MarshalECDSASignature(r, s)
}

// IsLowS 签名 s 大于椭圆曲线 base point 的阶的一半时返回 true。
func IsLowS(k *ecdsa.PublicKey, s *big.Int) (bool, error) {
	halfOrder, ok := curveHalfOrders[k.Curve]
	if !ok {
		return false, fmt.Errorf("curve not recognized [%s]", k.Curve)
	}

	return s.Cmp(halfOrder) != 1, nil
}

// ToLowS 如果签名 s 大于椭圆曲线 base point 的阶的一半，那么就用 base point 的阶减去 s，
// 并让 s 等于它。这样的转化不会破坏签名的正确性。
func ToLowS(k *ecdsa.PublicKey, s *big.Int) (*big.Int, error) {
	lowS, err := IsLowS(k, s)
	if err != nil {
		return nil, err
	}

	if !lowS {
		s.Sub(k.Params().N, s)
		return s, nil
	}

	return s, nil
}
