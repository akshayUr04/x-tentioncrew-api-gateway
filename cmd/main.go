package main

import (
	"log"
	"x-tentioncrew/api-gateway/pkg/config"
	"x-tentioncrew/api-gateway/pkg/user"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		log.Fatalln()
	}
	r := gin.Default()
	user.RegisterRoutes(r, &cfg)
	r.Run(cfg.Port)
}
