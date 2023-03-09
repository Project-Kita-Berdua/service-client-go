package product

import (
	"github.com/gin-gonic/gin"
	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/config"
	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/auth"
	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/product/routes"
	product "github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/product/pb"
	"google.golang.org/grpc"
	"fmt"
)

type ServiceClient struct {
	Client product.ProductServiceClient
}

func InitServiceClient(c *config.Config) product.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Couldn't connect: ", err)
	}
	return product.NewProductServiceClient(cc)
}

func RegisterRoutes(r *gin.Engine,  c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.POST("/:id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}