package order

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/auth"
	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/config"
	order "github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/order/pb"
	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/order/routes"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client order.OrderServiceClient
}

func InitServiceClient(c *config.Config) order.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Couldn't connect: ", err)
	}
	return order.NewOrderServiceClient(cc)
}

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
