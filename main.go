package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"os"
	"zj-admin/cache"
	"zj-admin/config"
	"zj-admin/db"
	"zj-admin/router"
	"zj-admin/util"
)

func main() {
	config.Init()

	util.InitLog()

	if config.Debug() {
		util.PProf()
	}

	cache.Init()

	db := db.Init()

	defer db.Close()

	log.Info().Str("version", os.Getenv("version")).Interface("env", os.Getenv("ENV")).
		Bool("debug", config.Debug()).Str("host", config.Host()).Msg(config.ProjectName() + " running")

	r := gin.Default()
	// 加载路由
	router.Init(r)

	log.Fatal().Err(r.Run(":3001"))
}
