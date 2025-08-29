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
		user := v1.Group("/users")
		{
			userHandlerV1 := v1handler.NewUserHandler()
			user.GET("", userHandlerV1.GetUsersV1)
			user.GET("/:id", userHandlerV1.GetUsersByIdV1)
			user.POST("", userHandlerV1.PostUsersV1)
			user.PUT("/:id", userHandlerV1.PutUsersV1)
			user.DELETE("/:id", userHandlerV1.DeleteUsersV1)
		}

		product := v1.Group("/products")
		{
			productHandlerV1 := v1handler.NewProductHandler()
			product.GET("", productHandlerV1.GetProductV1)
			product.GET("/:id", productHandlerV1.GetProductByIdV1)
			product.POST("", productHandlerV1.PostProductV1)
			product.PUT("/:id", productHandlerV1.PutProductV1)
			product.DELETE("/:id", productHandlerV1.DeleteProductV1)
		}
	}

	v2 := r.Group("/api/v2")
	{
		user := v2.Group("/users")
		{
			userHandlerV2 := v2handler.NewUserHandler()
			user.GET("", userHandlerV2.GetUsersV2)
			user.GET("/:id", userHandlerV2.GetUsersByIdV2)
			user.POST("", userHandlerV2.PostUsersV2)
			user.PUT("/:id", userHandlerV2.PutUsersV2)
			user.DELETE("/:id", userHandlerV2.DeleteUsersV2)
		}

	}

	r.Run(":8080")
}
