package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"app/api"
	"app/config"
	"app/pkg/db"
)

func main() {

	cfg := config.Load()

	db, err := db.NewConnectPostgres(cfg)
	if err != nil {
		log.Fatal("failed connection database: ", err.Error())
	}

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	api.NewApi(r, db)

	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}
