package service2

import (
	"fmt"
	"x-tentioncrew/api-gateway/pkg/config"
	"x-tentioncrew/api-gateway/pkg/service2/routes"

	"github.com/gin-gonic/gin"
)

func RegistrSvc2Routes(r *gin.Engine, cfg *config.Config) {
	svc := Svc2Client{
		Client: InitSvc2(cfg),
	}
	fmt.Println("in resister-----")
	routes := r.Group("/svc")
	routes.GET("/methods", svc.methods)
}

func (s *Svc2Client) methods(ctx *gin.Context) {
	fmt.Println("in r1-----")
	routes.Methods(ctx, s.Client)
}
