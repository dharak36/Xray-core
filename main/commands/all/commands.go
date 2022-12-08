package all

import (
	"github.com/dharak36/xray-core/main/commands/all/api"
	"github.com/dharak36/xray-core/main/commands/all/tls"
	"github.com/dharak36/xray-core/main/commands/base"
)

// go:generate go run github.com/dharak36/xray-core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
	)
}
