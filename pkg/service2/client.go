package service2

import (
	"fmt"
	"x-tentioncrew/api-gateway/pkg/config"
	"x-tentioncrew/api-gateway/pkg/service2/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Svc2Client struct {
	Client pb.Service2Client
}

func InitSvc2(c *config.Config) pb.Service2Client {
	fmt.Println("in init-----")
	cc, err := grpc.Dial(c.SERVICE2URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewService2Client(cc)
}
