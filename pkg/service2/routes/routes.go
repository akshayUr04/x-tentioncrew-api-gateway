package routes

import (
	"context"
	"fmt"
	"net/http"
	"x-tentioncrew/api-gateway/pkg/models"
	"x-tentioncrew/api-gateway/pkg/service2/pb"

	"github.com/gin-gonic/gin"
)

func Methods(ctx *gin.Context, c pb.Service2Client) {
	fmt.Println("in method-----")
	var data models.Data
	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "can't bind data",
		})
		return
	}
	fmt.Println("--------------")
	res, rpcErrr := c.Methods(context.Background(), &pb.MethodRequest{
		Method:   int64(data.Method),
		WaitTime: int64(data.Time),
	})

	if rpcErrr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   rpcErrr.Error(),
			"message": "can't getmethod data",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   &res,
	})

}
