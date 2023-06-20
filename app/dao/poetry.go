package dao

import (
	"windsmeller/app/logger"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetPoetryInfo(PoetryId int, TableName string, ctx gctx.Ctx) gdb.Record {
	sqlStr := "select title,author,paragraphs from " + TableName + " where id = ? ;"
	PoetryInfo, err := g.DB().GetOne(ctx, sqlStr, PoetryId)
	if err != nil {
		logger.LogError("get id "+gconv.String(PoetryId)+" from "+TableName+" err.", ctx)
	}
	return PoetryInfo

}

func GetPoetryRandom(TableName string, ctx gctx.Ctx) gdb.Record {
	sqlStr := "select id,title,author,paragraphs from " + TableName + " ORDER BY RANDOM() LIMIT 1;"
	PoetryInfo, err := g.DB().GetOne(ctx, sqlStr)
	if err != nil {
		logger.LogError("get random from "+TableName+" err.", ctx)
		logger.LogError(gconv.String(err), ctx)
	}
	return PoetryInfo

}

func PoetryStar(OpenId string, ctx g.Ctx) gdb.Result {
	sqlStr := `
		SELECT 
		    s.code,
		    s.poetryid,
		    CASE s.code
		         
		        WHEN 1 THEN tsj.title
		       
		        WHEN 2 THEN scj.title
		       
		        ELSE ssj.title
		    END AS title
		FROM
		    star s
		         
		        LEFT JOIN
		    tangshi_jian tsj ON s.poetryid = tsj.id AND s.code = 1
		  
		        LEFT JOIN
		    songci_jian scj ON s.poetryid = scj.id AND s.code = 2
		      
		        LEFT JOIN
		    songshi_jian ssj ON s.poetryid = ssj.id AND s.code = 3
		WHERE
		    s.openid = ?
	`
	PoetryStar, err := g.DB().GetAll(ctx, sqlStr, OpenId)
	if err != nil {
		logger.LogError("get PoetryStar with openid as "+OpenId+" err.", ctx)
	}
	return PoetryStar
}

func PoetryErr(UserOpenId string, PoetryId, PoetryCode int, ctx g.Ctx) {
	_, err := g.DB().Insert(ctx, "errlog", gdb.Map{
		"code":     PoetryCode,
		"openid":   UserOpenId,
		"poetryid": PoetryId,
		"status":   1,
	})
	if err != nil {
		logger.LogError(gconv.String(err), ctx)
	}
}

func PoetryStarPoetry(UserOpenId string, PoetryId, PoetryCode int, ctx g.Ctx) {
	_, err := g.DB().Insert(ctx, "star", gdb.Map{
		"code":     PoetryCode,
		"openid":   UserOpenId,
		"poetryid": PoetryId,
	})
	if err != nil {
		logger.LogError(gconv.String(err), ctx)
	}
}
func PoetryStarDelete(UserOpenId string, PoetryId, PoetryCode int, ctx g.Ctx) {
	_, err := g.DB().Delete(ctx, "star", gdb.Map{
		"openid":   UserOpenId,
		"code":     PoetryCode,
		"poetryid": PoetryId,
	})
	if err != nil {
		logger.LogError(gconv.String(err), ctx)
	}
}
