package lesson

import (
	v1handler "github.com/fish171204/gin-framework/internal/api/v1/handler"
	v2handler "github.com/fish171204/gin-framework/internal/api/v2/handler"
	"github.com/gin-gonic/gin"
)

func Lesson03RouteGroup() {
	r := gin.Default()

	userV1Handler := v1handler.NewUserHandler()
	r.GET("/api/v1/users", userV1Handler.GetUsersV1)
	r.GET("/api/v1/users/:id", userV1Handler.GetUsersByIdV1)
	r.POST("/api/v1/users", userV1Handler.PostUsersV1)
	r.PUT("/api/v1/users/:id", userV1Handler.PutUsersV1)
	r.DELETE("api/v1/users/:id", userV1Handler.DeleteUsersV1)

	productV1Handler := v1handler.NewProductHandler()
	r.GET("/api/v1/products", productV1Handler.GetProductV1)
	r.GET("/api/v1/products/:id", productV1Handler.GetProductByIdV1)
	r.POST("/api/v1/products", productV1Handler.PostProductV1)
	r.PUT("/api/v1/products/:id", productV1Handler.PutProductV1)
	r.DELETE("api/v1/products/:id", productV1Handler.DeleteProductV1)

	userV2Handler := v2handler.NewUserHandler()
	r.GET("/api/v2/users", userV2Handler.GetUsersV2)
	r.GET("/api/v2/users/:id", userV2Handler.GetUsersByIdV2)
	r.POST("/api/v2/users", userV2Handler.PostUsersV2)
	r.PUT("/api/v2/users/:id", userV2Handler.PutUsersV2)
	r.DELETE("api/v2/users/:id", userV2Handler.DeleteUsersV2)

	r.Run(":8080")
}
