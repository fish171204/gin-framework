package lesson

import (
	v1handler "github.com/fish171204/gin-framework/internal/api/v1/handler"
	v2handler "github.com/fish171204/gin-framework/internal/api/v2/handler"
	"github.com/gin-gonic/gin"
)

func Lesson03RouteGroup() {
	r := gin.Default()

	v1 := r.Group("/api/v1/")
	{
		userV1Handler := v1handler.NewUserHandler()
		v1.GET("/users", userV1Handler.GetUsersV1)
		v1.GET("/users/:id", userV1Handler.GetUsersByIdV1)
		v1.POST("/users", userV1Handler.PostUsersV1)
		v1.PUT("/users/:id", userV1Handler.PutUsersV1)
		v1.DELETE("/users/:id", userV1Handler.DeleteUsersV1)

		productV1Handler := v1handler.NewProductHandler()
		v1.GET("/products", productV1Handler.GetProductV1)
		v1.GET("/products/:id", productV1Handler.GetProductByIdV1)
		v1.POST("/products", productV1Handler.PostProductV1)
		v1.PUT("/products/:id", productV1Handler.PutProductV1)
		v1.DELETE("/products/:id", productV1Handler.DeleteProductV1)
	}

	v2 := r.Group("/api/v2")
	{
		userV2Handler := v2handler.NewUserHandler()
		v2.GET("/users", userV2Handler.GetUsersV2)
		v2.GET("/users/:id", userV2Handler.GetUsersByIdV2)
		v2.POST("/users", userV2Handler.PostUsersV2)
		v2.PUT("/users/:id", userV2Handler.PutUsersV2)
		v2.DELETE("/users/:id", userV2Handler.DeleteUsersV2)
	}

	r.Run(":8080")
}
