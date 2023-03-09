package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/auth"
	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/config"
	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/order"
	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/product"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	// h := config.Init(c.DBURL)

	// jwt := utils.JwtWrapper{
	// 	SecretKey:       c.JwtSecretKey,
	// 	Issuer:          "go-grpc-auth-svc",
	// 	ExpirationHours: 24 * 365,
	// }

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	// lis, err := net.Listen("tcp", c.AuthSvcUrl) // port of user service
	// if err != nil {
	// 	log.Fatalln("Failed to listening: ", err)
	// }

	// fmt.Println("Auth Svc on : ", c.AuthSvcUrl)

	// s := services.AuthService{
	// 	H:   h,
	// 	Jwt: jwt,
	// }

	// grpcServer := grpc.NewServer()

	// pb.RegisterAuthServiceServer(grpcServer, &s)

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalln("Failed to serve: ", err)
	// }

	r.Run(c.Port)
}
