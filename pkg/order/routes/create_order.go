package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/go-grpc-http-rest-microservice-tutorial/pkg/order/pb"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
	// UserId    int64 `json:"user_id"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
