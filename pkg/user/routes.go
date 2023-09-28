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
	routes.GET("/:id", svc.GetUser)
	routes.PATCH("/:id", svc.UpdateUser)
	routes.DELETE("/:id", svc.DeleteUser)
}

func (svc *UserSvcClient) CreateUser(ctx *gin.Context) {
	routes.CreateUser(ctx, svc.Client)
}

func (svc *UserSvcClient) GetUser(ctx *gin.Context) {
	routes.GetUserById(ctx, svc.Client)
}

func (svc *UserSvcClient) UpdateUser(ctx *gin.Context) {
	routes.UpdateUser(ctx, svc.Client)
}

func (svc *UserSvcClient) DeleteUser(ctx *gin.Context) {
	routes.DeleteUser(ctx, svc.Client)
}
