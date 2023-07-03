package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCurveHalf(t *testing.T) {
	t.Log(curveHalfOrders[elliptic.P224()].String())
	t.Log(elliptic.P224().Params().N)

	t.Log(curveHalfOrders[elliptic.P256()].String())
	t.Log(elliptic.P256().Params().N)

	t.Log(curveHalfOrders[elliptic.P384()].String())
	t.Log(elliptic.P384().Params().N)

	t.Log(curveHalfOrders[elliptic.P521()].String())
	t.Log(elliptic.P521().Params().N)
}

func TestHalfCurveSignature(t *testing.T) {
	sk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	for i := 0; ; i++ {
		m := fmt.Sprintf("message: %d", i)
		h := sha256.Sum256([]byte(m))
		sig, err := sk.Sign(rand.Reader, h[:], nil)
		require.NoError(t, err)
		r, s, err := UnmarshalECDSASignature(sig)
		if err != nil {
			t.Log(err)
			return
		}
		isLowS, err := IsLowS(&sk.PublicKey, s)
		require.NoError(t, err)

		if !isLowS {
			t.Log("found...")
			t.Log("origin s:", s.String())
			s, err = ToLowS(&sk.PublicKey, s)
			t.Log("after s:", s.String())
			require.NoError(t, err)
			ok := ecdsa.Verify(&sk.PublicKey, h[:], r, s)
			require.Equal(t, ok, true)
			break
		}
	}
}

func TestUnmarshalECDSASignature(t *testing.T) {
	_, _, err := UnmarshalECDSASignature(nil)
	require.Contains(t, err.Error(), "failed to unmarshal signature [")

	_, _, err = UnmarshalECDSASignature([]byte{})
	require.Contains(t, err.Error(), "failed to unmarshal signature [")

	_, _, err = UnmarshalECDSASignature([]byte{0})
	require.Contains(t, err.Error(), "failed to unmarshal signature [")

	sig, err := MarshalECDSASignature(big.NewInt(-1), big.NewInt(1))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Equal(t, err.Error(), "invalid signature, R must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(1), big.NewInt(-1))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Equal(t, err.Error(), "invalid signature, S must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(-1), big.NewInt(-1))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Equal(t, err.Error(), "invalid signature, R must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(0), big.NewInt(1))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Equal(t, err.Error(), "invalid signature, R must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(1), big.NewInt(0))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Equal(t, err.Error(), "invalid signature, S must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(0), big.NewInt(0))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Equal(t, err.Error(), "invalid signature, R must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(1), big.NewInt(1))
	require.NoError(t, err)
	r, s, err := UnmarshalECDSASignature(sig)
	require.NoError(t, err)
	require.Equal(t, int64(1), r.Int64())
	require.Equal(t, int64(1), s.Int64())
}
