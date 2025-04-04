package main

import (
	"store/src/configs"
	"store/src/servers"
)

func main() {
	cfg := configs.GetConfig()
	servers.NewServer(*cfg)
}