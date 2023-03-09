package auth

import (
	"fmt"

	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/auth/pb"
	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Couldn't connect: ", err)
	}

	return pb.NewAuthServiceClient(cc)
}
