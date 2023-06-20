package dao

import (
	"windsmeller/app/logger"
	"windsmeller/app/model"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

//

func VerifyOpenIdExist(openId string, ctx g.Ctx) bool {

	// if openId is in db,return true.else return false
	var (
		key = openId
	)
	sqlStr := "select count(1) as a from user where openid = ? ;"
	userExistCount, err := g.DB().GetOne(ctx, sqlStr, key)

	if err != nil {
		logger.LogError(gconv.String(err), ctx)
		return true
	} else {
		openIdExist := gconv.Int(userExistCount["a"])

		if openIdExist == 0 {
			return false
		} else if openIdExist == 1 {
			return true
		} else {
			logger.LogWarn(openId+" is more than 1,check it.", ctx)
			logger.LogInfo(openId+" is more than 1,check it.", ctx)
			return true
		}
	}

}

func VerifyUserExist(userName string, ctx g.Ctx) bool {

	// if username is in db,return true.else return false
	var (
		key = userName
	)
	sqlStr := "select count(1) as a from user where user = ? ;"
	userExistCount, err := g.DB().GetOne(ctx, sqlStr, key)

	if err != nil {
		logger.LogError(gconv.String(err), ctx)
		return true
	} else {
		userExist := gconv.Int(userExistCount["a"])

		if userExist == 0 {
			return false
		} else if userExist == 1 {
			return true
		} else {
			logger.LogWarn(userName+" is more than 1,check it.", ctx)
			logger.LogInfo(userName+" is more than 1,check it.", ctx)
			return true
		}
	}

}

func VerifyUser(UserName, UserPass string, ctx g.Ctx) bool {
	// VerifyUser name and pass.same return true,else return fasle
	sqlStr := "select count(1) as VerifyUserStatus from user where user = ? and pass = ? ;"
	verifyUserStatus, err := g.DB().GetOne(ctx, sqlStr, UserName, UserPass)
	if err != nil {
		logger.LogError("VerifyUser err,UserName is "+gconv.String(UserName), ctx)
		return false
	} else {
		VerifyUserStatus := gconv.Int(verifyUserStatus["VerifyUserStatus"])
		if VerifyUserStatus == 0 {
			return false

		} else {
			return true
		}
	}

}

func UserAccessTimeRefresh(openId string, ctx g.Ctx) {
	_, err := g.DB().Update(ctx, "user", "lastaccesstime=CURRENT_TIMESTAMP", "openId=?", openId)
	if err != nil {
		logger.LogError("Refresh user last Access Time  err,Uid is "+openId, ctx)

	}
}

func GetUserName(openId string, ctx g.Ctx) string {
	sqlStr := "select user from user where openid = ?"
	userInfo, err := g.DB().GetOne(ctx, sqlStr, openId)
	if err != nil {
		logger.LogError("get user name err,openId is "+openId, ctx)

		return ""
	} else {
		UserName := gconv.String(userInfo["user"])
		return UserName
	}

}

func GetUserId(UserName string, ctx g.Ctx) int {
	sqlStr := "select id from user where user = ?"
	userInfo, err := g.DB().GetOne(ctx, sqlStr, UserName)
	if err != nil {
		logger.LogError("get user info err,UserName is "+gconv.String(UserName), ctx)

		return 0
	} else {
		UserId := gconv.Int(userInfo["id"])
		return UserId
	}

}

func GetUserInfo(userOpenId string, ctx g.Ctx) model.UserInfoWithOpenId {

	sqlStr := "select * from user where openid = ?"
	userInfo, err := g.DB().GetOne(ctx, sqlStr, userOpenId)

	if err != nil {
		logger.LogError("get user info err,user openid is "+userOpenId, ctx)
		UserInfo := model.UserInfoWithOpenId{}
		return UserInfo
	} else {
		UserInfo := model.UserInfoWithOpenId{
			UserOpenId: userOpenId,
			UserName:   gconv.String(userInfo["user"]),
		}
		return UserInfo
	}
}

func RegisterUser(UserInfo model.UserInfoWithOpenId, ctx g.Ctx) model.UserInfoWithOpenId {

	_, err := g.DB().Insert(ctx, "user", gdb.Map{
		"user":   UserInfo.UserName,
		"openid": UserInfo.UserOpenId,
	})
	if err != nil {
		logger.LogError(gconv.String(err), ctx)
		NilUser := model.UserInfoWithOpenId{}
		return NilUser
	} else {

		UserInfo = GetUserInfo(UserInfo.UserOpenId, ctx)
		return UserInfo
	}
}
