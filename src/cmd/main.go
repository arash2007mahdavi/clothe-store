package main

import (
	"store/src/cache"
	"store/src/configs"
	"store/src/servers"
)

func main() {
	cfg := configs.GetConfig()
	cache.InitRedis()
	defer cache.CloseRedis()
	servers.NewServer(*cfg)
}