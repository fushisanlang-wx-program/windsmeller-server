package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	ctx = gctx.New()
)

func returnErrCode(r *ghttp.Request, code int, msg string) {
	r.Response.Status = code
	r.Response.WriteJson(g.Map{
		"message": msg,
	})
}
