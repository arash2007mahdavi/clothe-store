package main

import (
	"store/src/loggers"
	"store/src/cache"
	"store/src/configs"
	"store/src/servers"
)

func main() {
	cfg := configs.GetConfig()
	logger := loggers.NewLogger(cfg)
	cache.InitRedis()
	defer cache.CloseRedis()
	logger.Info(loggers.General, loggers.Startup, "server started", nil)
	servers.NewServer(*cfg)
}