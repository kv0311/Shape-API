package main

import (
	"shape-api/config"
	"shape-api/db"
	"shape-api/logger"
	"shape-api/server"
)

func main() {
	//Load the .env file
	config.Init()
	err := db.InitPG()
	if err != nil {
		logger.Logger.Error("Fail to connect redis: ", err)
	}
	db.InitRedis()
	server.Init()
}
