package routes

import (
	"context"
	"net/http"
	"strconv"
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
	res, err := p.CreateUser(context.Background(), &pb.CreatUserRequest{
		Name:        userDetails.Name,
		HouseName:   userDetails.HouseName,
		City:        userDetails.City,
		Email:       userDetails.Email,
		PhoneNumber: userDetails.PhoneNumber,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"rpcErr":  &res,
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

func GetUserById(ctx *gin.Context, p pb.UserSvcClient) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errot":   err.Error(),
			"message": "cant find userId",
		})
		return
	}

	result, err := p.GetUserById(ctx, &pb.GetUserRequest{
		Id: id,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "can't make request to delete user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user details",
		"data":    &result,
	})
}

func UpdateUser(ctx *gin.Context, p pb.UserSvcClient) {
	var userDetails models.UserDetails
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errot":   err.Error(),
			"message": "cant find userId",
		})
		return
	}

	err = ctx.BindJSON(&userDetails)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "can't bind data",
		})
		return
	}
	result, err := p.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:          id,
		Name:        userDetails.Name,
		HouseName:   userDetails.HouseName,
		City:        userDetails.City,
		Email:       userDetails.Email,
		PhoneNumber: userDetails.PhoneNumber,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"rpcErr":  &result,
			"error":   err.Error(),
			"message": "can't make request to update user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user updated",
		"data":    &result,
	})
}

func DeleteUser(ctx *gin.Context, p pb.UserSvcClient) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errot":   err.Error(),
			"message": "cant find userId",
		})
		return
	}

	res, err := p.DeleteUser(ctx, &pb.DeleteUserRequest{
		Id: id,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"rpcErr":  &res,
			"error":   err.Error(),
			"message": "can't make request to delete user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user deleted",
		"data":    &res,
	})
}
