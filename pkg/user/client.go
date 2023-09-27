package user

import (
	"fmt"
	"x-tentioncrew/api-gateway/pkg/config"
	"x-tentioncrew/api-gateway/pkg/user/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserSvcClient struct {
	Client pb.UserSvcClient
}

func InitUserSvc(c *config.Config) pb.UserSvcClient {
	cc, err := grpc.Dial(c.UserSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewUserSvcClient(cc)
}
