package routes

import (
	"context"
	"net/http"
	"x-tentioncrew/api-gateway/pkg/models"
	"x-tentioncrew/api-gateway/pkg/service2/pb"

	"github.com/gin-gonic/gin"
)

func Methods(ctx *gin.Context, c pb.ServiceClient) {
	var data models.Data
	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "can't bind data",
		})
		return
	}
	res, err := c.Methods(context.Background(), &pb.MethodRequest{
		Method:   int64(data.Method),
		WaitTime: int64(data.Time),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   &res,
	})

}
