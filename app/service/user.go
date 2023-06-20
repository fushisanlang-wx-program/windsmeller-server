package service

import (
	"bytes"
	"net/http"
	"windsmeller/app/dao"
	"windsmeller/app/logger"
	"windsmeller/app/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func VerifyUserExist(userName string, ctx g.Ctx) bool {
	userExistStatus := dao.VerifyUserExist(userName, ctx)
	return userExistStatus
}

func VerifyOpenIdExist(OpenId string, ctx g.Ctx) bool {
	openIdExistStatus := dao.VerifyOpenIdExist(OpenId, ctx)
	return openIdExistStatus
}
func VerifyUser(UserName, UserPass string, ctx g.Ctx) bool {
	VerifyUserStatus := dao.VerifyUser(UserName, UserPass, ctx)
	return VerifyUserStatus
}

func UserAccessTimeRefresh(openId string, ctx g.Ctx) {
	dao.UserAccessTimeRefresh(openId, ctx)
}

func GetUserId(UserName string, ctx g.Ctx) int {
	UserId := dao.GetUserId(UserName, ctx)
	return UserId
}
func GetUserName(OpenId string, ctx g.Ctx) string {
	UserName := dao.GetUserName(OpenId, ctx)
	return UserName
}
func RegisterUser(userName, userOpenId string, ctx g.Ctx) model.UserInfoWithOpenId {
	UserInfo := model.UserInfoWithOpenId{
		//UserId:             userId,
		UserName:   userName,
		UserOpenId: userOpenId,
		UserAdmin:  false,
		//UserCreateTime:     "",
		//UserLastAccessTime: "",
	}
	UserInfo = dao.RegisterUser(UserInfo, ctx)
	return UserInfo
}
func GetOpenId(CodeStr string, ctx g.Ctx) model.OpenId {
	logger.LogInfo("get codestr as "+CodeStr, ctx)
	GetStatus, OpenId := getOpenId(CodeStr, ctx)
	var code int
	var message, userName string
	if GetStatus == true {

		openIdExist := VerifyOpenIdExist(OpenId, ctx)

		if openIdExist == true {
			UserAccessTimeRefresh(OpenId, ctx)
			userName = GetUserName(OpenId, ctx)
			//用户存在
			code = 200
			message = ""
		} else {
			//用户不存在
			code = 401
			message = "用户未注册"
			userName = ""
		}

	} else {
		code = 502
		message = "微信服务器异常，请重试或联系开发者"
		userName = ""
	}
	var UserOpenId = model.OpenId{
		Code:     code,
		OpenId:   OpenId,
		Message:  message,
		UserName: userName,
	}
	logger.LogInfo("get user openid as "+OpenId, ctx)
	return UserOpenId
}

func getOpenId(codeStr string, ctx g.Ctx) (bool, string) {
	jscode2sessionUrlPath, err := g.Cfg().Get(ctx, "wx.jscode2sessionUrlPath")
	if err != nil {
		logger.LogError("urlPath not config.", ctx)
	}
	appid, err := g.Cfg().Get(ctx, "wx.appid")
	if err != nil {
		logger.LogError("appid not config.", ctx)
	}
	secret, err := g.Cfg().Get(ctx, "wx.secret")
	if err != nil {
		logger.LogError("secret not config.", ctx)
	}
	grantType, err := g.Cfg().Get(ctx, "wx.grant_type")
	if err != nil {
		logger.LogError("grant_type not config.", ctx)
	}
	urlPath := gconv.String(jscode2sessionUrlPath)
	postString := "appid=" + gconv.String(appid) + "&secret=" + gconv.String(secret) + "&js_code=" + gconv.String(codeStr) + "&grant_type=" + gconv.String(grantType)
	postStringByte := []byte(postString)
	req, err := http.NewRequest("Get", urlPath, bytes.NewBuffer(postStringByte))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.LogError(gconv.String(err), ctx)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		OpenId := transformation(resp)["openid"]
		return true, gconv.String(OpenId)
	} else {
		return false, ""
	}

}

func VerifyCode(r *ghttp.Request, ctx g.Ctx) model.OpenId {
	CodeStr := r.Get("code").String()
	SessioStatus := GetOpenId(CodeStr, ctx)
	return SessioStatus
}

func VerifySession(r *ghttp.Request, ctx g.Ctx) (bool, string, string) {

	sessionData, err := r.Session.Data()
	if err != nil {
		return false, "", ""
	}
	var userStruct *model.UserInfoWithOpenId

	if gconv.Struct(sessionData, &userStruct) != nil {
		return false, "", ""
	}
	UserName := userStruct.UserName

	UserOpenId := userStruct.UserOpenId
	userNameTrue := GetUserName(UserOpenId, ctx)
	if UserOpenId == "" {
		return false, "", ""
	} else if UserName == userNameTrue {
		return true, UserOpenId, UserName
	} else {
		return false, "", ""
	}

}
