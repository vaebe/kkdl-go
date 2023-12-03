package main

import (
	_ "compressURL/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"compressURL/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
