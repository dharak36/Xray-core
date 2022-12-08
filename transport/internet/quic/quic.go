package quic

import (
	"github.com/dharak36/xray-core/common"
	"github.com/dharak36/xray-core/transport/internet"
)

//go:generate go run github.com/dharak36/xray-core/common/errors/errorgen

// Here is some modification needs to be done before update quic vendor.
// * use bytespool in buffer_pool.go
// * set MaxReceivePacketSize to 1452 - 32 (16 bytes auth, 16 bytes head)
//
//

const (
	protocolName   = "quic"
	internalDomain = "quic.internal.example.com"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
