package user

import (
	"x-tentioncrew/api-gateway/pkg/config"
	"x-tentioncrew/api-gateway/pkg/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	svc := &UserSvcClient{
		Client: InitUserSvc(cfg),
	}

	routes := r.Group("/user")
	routes.POST("/create", svc.CreateUser)
}

func (svc *UserSvcClient) CreateUser(ctx *gin.Context) {
	routes.CreateUser(ctx, svc.Client)
}
