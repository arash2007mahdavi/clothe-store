package main

import (
	"store/src/configs"
	"store/src/database"
	"store/src/database/migrations"
	"store/src/loggers"
	"store/src/servers"
)

func main() {
	cfg := configs.GetConfig()
	logger := loggers.NewLogger(cfg)
	database.InitRedis()
	defer database.CloseRedis()
	err := database.InitDB(cfg)
	if err != nil {
		logger.Fatal(loggers.Postgres, loggers.Startup, "failed in start database", nil)
		panic(err)
	}
	defer database.CloseDB()
	migrations.Up_1()
	logger.Info(loggers.General, loggers.Startup, "server started", nil)
	servers.NewServer(*cfg)
}