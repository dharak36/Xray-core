package tcp

import (
	"github.com/dharak36/xray-core/common"
	"github.com/dharak36/xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
