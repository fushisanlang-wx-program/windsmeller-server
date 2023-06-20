package app

import (
	"time"
	"windsmeller/app/api"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gsession"
)

func Run() {

	//配置文件
	//g.Cfg().SetFileName("config.yml")
	//服务定义
	s := g.Server()
	//session相关
	s.SetSessionMaxAge(0 * time.Hour)
	//s.SetSessionStorage(gsession.NewStorageRedis(g.Redis("session")))
	//s.SetSessionStorage(gsession.NewStorageMemory())
	s.SetSessionStorage(gsession.NewStorageRedis(g.Redis()))
	//s.BindHandler("/", api.RootPage)
	//group := s.Group("/status")
	//服务器版本
	//group.ALL("/version", api.GetVersion)
	group := s.Group("/wind/sys")
	group.GET("/info", api.SysInfo)
	group = s.Group("/wind/user")
	group.GET("/getopenid", api.GetOpenId)
	group.GET("/register", api.UserRegister)
	//group.GET("/signin", api.UserSignIn)
	group = s.Group("/wind/poetry")
	group.GET("/read/:codeid/:id", api.PoetryRead)
	group.GET("/star", api.PoetryStar)
	group.DELETE("/star", api.PoetryStarDelete)
	group.POST("/star", api.PoetryStarPoetry)
	group.GET("/random", api.PoetryRandom)
	group.GET("/err", api.PoetryErr)
	s.Run()
}
