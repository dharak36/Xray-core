package udp

import (
	"github.com/dharak36/xray-core/common"
	"github.com/dharak36/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
