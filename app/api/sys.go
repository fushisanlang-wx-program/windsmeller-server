package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func SysInfo(r *ghttp.Request) {
	r.Response.WriteJson(g.Map{
		"info": "windsmeller",
	})
}
