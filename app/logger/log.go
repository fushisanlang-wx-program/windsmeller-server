package logger

import (
	"github.com/gogf/gf/v2/frame/g"
)

func LogInfo(logString string, ctx g.Ctx) {
	g.Log().Info(ctx, logString)
}
func LogWarn(logString string, ctx g.Ctx) {
	g.Log().Warning(ctx, logString)
}
func LogError(logString string, ctx g.Ctx) {
	g.Log().Error(ctx, logString)
}
