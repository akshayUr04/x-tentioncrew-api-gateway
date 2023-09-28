package service2

import (
	"x-tentioncrew/api-gateway/pkg/config"
	"x-tentioncrew/api-gateway/pkg/service2/routes"

	"github.com/gin-gonic/gin"
)

func RegistrSvc2Routes(r *gin.Engine, cfg *config.Config) {
	svc := Svc2Client{
		Client: InitSvc2(cfg),
	}
	routes := r.Group("/svc")
	routes.GET("/methods", svc.methods)
}

func (s *Svc2Client) methods(ctx *gin.Context) {
	routes.Methods(ctx, s.Client)
}
