package upgrade

import "github.com/zeromicro/go-zero/tools/goctl/internal/cobrax"

var Cmd = cobrax.NewCommand("upgrade", cobrax.WithRunE(upgrade))
