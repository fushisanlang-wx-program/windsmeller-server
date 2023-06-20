package api

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"windsmeller/app/service"
)

func PoetryErr(r *ghttp.Request) {

	verifyStatus := service.VerifyCode(r, ctx)

	if verifyStatus.Code == 200 {
		UserOpenId := verifyStatus.OpenId
		PoetryId := r.Get("poetryid").Int()
		PoetryCode := r.Get("poetrycode").Int()
		service.PoetryErr(UserOpenId, PoetryId, PoetryCode, ctx)
		returnErrCode(r, 200, "谢谢")

	} else {
		returnErrCode(r, 401, "用户校验失败,请重新登录")
	}

}

func PoetryRandom(r *ghttp.Request) {

	verifyStatus := service.VerifyCode(r, ctx)

	if verifyStatus.Code == 200 {

		PoetryInfo := service.GetPoetryRandom(ctx)
		r.Response.WriteJson(PoetryInfo)

	} else {
		returnErrCode(r, 401, "用户校验失败,请重新登录")
	}

}

func PoetryRead(r *ghttp.Request) {

	id := r.Get("id")
	PoetryId := gconv.Int(id)
	code := r.Get("codeid")
	PoetryCode := gconv.Int(code)
	if PoetryId == 0 {
		returnErrCode(r, 417, "id err")
	} else if PoetryCode < 1 || PoetryCode > 3 {
		returnErrCode(r, 417, "code err")
	} else {

		PoetryInfo := service.GetPoetryInfo(PoetryId, PoetryCode, ctx)
		r.Response.WriteJson(PoetryInfo)
	}

}

func PoetryStar(r *ghttp.Request) {
	verifyStatus := service.VerifyCode(r, ctx)

	if verifyStatus.Code == 200 {

		PoetryStar := service.PoetryStar(verifyStatus.OpenId, ctx)
		r.Response.WriteJson(PoetryStar)

	} else {
		r.Response.WriteJson(verifyStatus)
	}

}
func PoetryStarPoetry(r *ghttp.Request) {
	verifyStatus := service.VerifyCode(r, ctx)

	if verifyStatus.Code == 200 {
		UserOpenId := verifyStatus.OpenId
		PoetryId := r.Get("poetryid").Int()
		PoetryCode := r.Get("poetrycode").Int()
		service.PoetryStarPoetry(UserOpenId, PoetryId, PoetryCode, ctx)
		returnErrCode(r, 200, "收藏成功")

	} else {
		r.Response.WriteJson(verifyStatus)
	}

}
func PoetryStarDelete(r *ghttp.Request) {
	verifyStatus := service.VerifyCode(r, ctx)

	if verifyStatus.Code == 200 {
		UserOpenId := verifyStatus.OpenId
		PoetryId := r.Get("poetryid").Int()
		PoetryCode := r.Get("poetrycode").Int()
		service.PoetryStarDelete(UserOpenId, PoetryId, PoetryCode, ctx)
		returnErrCode(r, 200, "删除成功")

	} else {
		r.Response.WriteJson(verifyStatus)
	}

}
