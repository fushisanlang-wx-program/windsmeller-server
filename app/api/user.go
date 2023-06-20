package api

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"windsmeller/app/logger"
	"windsmeller/app/service"
)

func GetOpenId(r *ghttp.Request) {

	CodeStr := r.Get("code").String()

	if CodeStr == "" {
		returnErrCode(r, 417, "数据空")
	} else {
		openId := service.GetOpenId(CodeStr, ctx)
		err := r.Session.Set("openId", openId.OpenId)
		if err != nil {
			logger.LogError("set session openid as "+openId.OpenId+" err!", ctx)
		}
		if openId.Code == 200 {
			err = r.Session.Set("userName", openId.UserName)
			if err != nil {
				logger.LogError("set session userName as "+openId.UserName+" err!", ctx)
			}
		}
		r.Response.WriteJson(openId)
	}
}

func UserRegister(r *ghttp.Request) {

	UserName := r.Get("UserName").String()
	UserCode := r.Get("code").String()
	// UserOpenId := r.Get("UserOpenId").String()
	UserOpenId := service.GetOpenId(UserCode, ctx).OpenId
	if UserName == "" || UserOpenId == "" {

		returnErrCode(r, 417, "用户注册失败，数据空")
	} else if service.VerifyOpenIdExist(UserOpenId, ctx) == true {
		returnErrCode(r, 423, "请勿重复注册")

	} else if service.VerifyUserExist(UserName, ctx) == true {
		returnErrCode(r, 423, "用户注册失败，用户名已存在")
	} else {
		UserInfo := service.RegisterUser(UserName, UserOpenId, ctx)
		err := r.Session.Set("userName", UserName)
		if err != nil {
			logger.LogError("set session userName as "+UserName+" err!", ctx)
		}
		r.Response.WriteJson(UserInfo)

	}
}

/*
func UserSignIn(r *ghttp.Request) {

	UserName := r.Get("UserName").String()
	UserPass := r.Get("UserPass").String()
	if UserName == "" || UserPass == "" {
		returnErrCode(r, 417, "用户登录失败，数据空")
	} else if service.VerifyUser(UserName, UserPass, ctx) == true {
		Uid := service.GetUserId(UserName, ctx)
		service.UserAccessTimeRefresh(Uid, ctx)
		err := r.Session.Set("userName", UserName)
		if err != nil {
			logger.LogError("get username err,username get "+UserName, ctx)
			returnErrCode(r, 401, "UserName 异常")
		}
		err = r.Session.Set("Uid", Uid)
		if err != nil {
			logger.LogError("get Uid err,Uid get "+gconv.String(Uid), ctx)
			returnErrCode(r, 401, "Uid 异常")
		}
		r.Response.WriteJson(g.Map{
			"message":  "用户登录成功",
			"userName": UserName,
		})
	} else {
		returnErrCode(r, 401, "用户登录失败,账户密码不匹配")

	}
}
*/
