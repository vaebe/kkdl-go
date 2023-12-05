package main

import (
	_ "compressURL/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "compressURL/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"compressURL/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
