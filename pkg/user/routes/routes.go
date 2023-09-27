package routes

import (
	"context"
	"net/http"
	"x-tentioncrew/api-gateway/pkg/models"
	"x-tentioncrew/api-gateway/pkg/user/pb"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context, p pb.UserSvcClient) {
	var userDetails models.UserDetails
	err := ctx.BindJSON(&userDetails)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "can't bind data",
		})
		return
	}
	res, err := p.CreatUser(context.Background(), &pb.CreatUserRequest{
		Name:        userDetails.Name,
		HouseName:   userDetails.HouseName,
		City:        userDetails.City,
		Email:       userDetails.Email,
		PhoneNumber: userDetails.PhoneNumber,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"rpcerr":  &res,
			"error":   err.Error(),
			"message": "can't make request",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created succesfully",
		"data":    &res,
	})

}
