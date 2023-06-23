/**
* Author: Xiangyu Wu
* Date: 2023-06-23
* From: hyperledger/fabric/common/metadata/metadata.go
 */

package metadata

// 这些数据可以通过 ldflags 进行传递。
var (
	Version         = "latest"
	CommitSHA       = "development build"
	BaseDockerLabel = "org.geistwelt.quarkx"
	DockerNamespace = "hyperledger"
)
