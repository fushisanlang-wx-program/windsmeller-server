package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"windsmeller/app"
	"windsmeller/app/logger"
)

func main() {
	var (
		ctx = gctx.New()
	)

	logger.LogInfo("服务启动", ctx)
	app.Run()
}
