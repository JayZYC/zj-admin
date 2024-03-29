package main

import (
	"os"
	"zj-admin/cache"
	"zj-admin/config"
	"zj-admin/db"
	"zj-admin/router"
	"zj-admin/util"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	config.Init()

	util.InitLog()

	if config.Debug() {
		util.PProf()
	}

	cache.Init()

	db.Init()

	log.Info().Str("version", os.Getenv("version")).Interface("env", os.Getenv("ENV")).
		Bool("debug", config.Debug()).Msg(config.ProjectName() + " running")

	r := gin.Default()
	// 加载路由
	router.Init(r)

	log.Fatal().Err(r.Run(":1010"))
}
