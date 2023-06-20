package service

import (
	"windsmeller/app/dao"
	"windsmeller/app/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

func GetPoetryInfo(PoetryId, PoetryCode int, ctx gctx.Ctx) model.PoetryInfo {
	var TableName string
	if PoetryCode == 1 {
		//tangshi——jian
		TableName = "tangshi_jian"

		//} else if PoetryCode == 2 {
		//	//tangshi——fan
		//	TableName = "tangshi"
	} else if PoetryCode == 2 {
		//songci——jian
		TableName = "songci_jian"
		//} else if PoetryCode == 4 {
		//	//songci——jian
		//	TableName = "songci"
	} else if PoetryCode == 3 {
		//songci——jian
		TableName = "songshi_jian"
		//} else {
		//	//songci——jian
		//	TableName = "songshi"
	}

	poetryInfo := dao.GetPoetryInfo(PoetryId, TableName, ctx)
	paragraphs := gconv.String(poetryInfo["paragraphs"])
	//去除字符串首的[' 和尾的 ']
	patternStr := `\['|'\]`
	result, _ := gregex.Replace(patternStr, []byte(""), []byte(paragraphs))

	// 切割成list
	patternStr = `', '`
	result2 := gregex.Split(patternStr, string(result))

	PoetryInfo := model.PoetryInfo{
		PoetryCode:       PoetryCode,
		PoetryId:         PoetryId,
		PoetryAuthor:     gconv.String(poetryInfo["author"]),
		PoetryTitle:      gconv.String(poetryInfo["title"]),
		PoetryParagraphs: result2,
	}
	return PoetryInfo
}

func PoetryStar(OpenId string, ctx g.Ctx) gdb.Result {
	PoetryStar := dao.PoetryStar(OpenId, ctx)
	return PoetryStar
}

func GetPoetryRandom(ctx gctx.Ctx) model.PoetryInfo {
	PoetryCode := grand.N(1, 3)
	var TableName string
	if PoetryCode == 1 {
		//tangshi——jian
		TableName = "tangshi_jian"

		//} else if PoetryCode == 2 {
		//	//tangshi——fan
		//	TableName = "tangshi"
	} else if PoetryCode == 2 {
		//songci——jian
		TableName = "songci_jian"
		//} else if PoetryCode == 4 {
		//	//songci——jian
		//	TableName = "songci"
	} else if PoetryCode == 3 {
		//songci——jian
		TableName = "songshi_jian"
		//} else {
		//	//songci——jian
		//	TableName = "songshi"
	}

	poetryInfo := dao.GetPoetryRandom(TableName, ctx)
	paragraphs := gconv.String(poetryInfo["paragraphs"])
	//去除字符串首的[' 和尾的 ']
	patternStr := `\['|'\]`
	result, _ := gregex.Replace(patternStr, []byte(""), []byte(paragraphs))

	// 切割成list
	patternStr = `', '`
	result2 := gregex.Split(patternStr, string(result))

	PoetryInfo := model.PoetryInfo{
		PoetryId:         gconv.Int(poetryInfo["id"]),
		PoetryAuthor:     gconv.String(poetryInfo["author"]),
		PoetryTitle:      gconv.String(poetryInfo["title"]),
		PoetryParagraphs: result2,
		PoetryCode:       PoetryCode,
	}
	return PoetryInfo
}

func PoetryErr(UserOpenId string, PoetryId, PoetryCode int, ctx g.Ctx) {
	dao.PoetryErr(UserOpenId, PoetryId, PoetryCode, ctx)

}
func PoetryStarPoetry(UserOpenId string, PoetryId, PoetryCode int, ctx g.Ctx) {
	dao.PoetryStarPoetry(UserOpenId, PoetryId, PoetryCode, ctx)

}
func PoetryStarDelete(UserOpenId string, PoetryId, PoetryCode int, ctx g.Ctx) {
	dao.PoetryStarDelete(UserOpenId, PoetryId, PoetryCode, ctx)

}
