package lesson

import (
	v1handler "github.com/fish171204/gin-framework/internal/api/v1/handler"
	v2handler "github.com/fish171204/gin-framework/internal/api/v2/handler"
	"github.com/fish171204/gin-framework/utils"
	"github.com/gin-gonic/gin"
)

func Lesson03RouteGroup() {
	r := gin.Default()

	if err := utils.RegisterValidators(); err != nil {
		// stop
		panic(err)
	}

	v1 := r.Group("/api/v1/")
	{
		user := v1.Group("/users")
		{
			userHandlerV1 := v1handler.NewUserHandler()
			user.GET("", userHandlerV1.GetUsersV1)
			user.GET("/:id", userHandlerV1.GetUsersByIdV1)
			user.GET("/admin/:uuid", userHandlerV1.GetUsersByUuidV1)
			user.POST("", userHandlerV1.PostUsersV1)
			user.PUT("/:id", userHandlerV1.PutUsersV1)
			user.DELETE("/:id", userHandlerV1.DeleteUsersV1)
		}

		product := v1.Group("/products")
		{
			productHandlerV1 := v1handler.NewProductHandler()
			product.GET("", productHandlerV1.GetProductV1)
			product.GET("/:slug", productHandlerV1.GetProductBySlugV1)
			product.POST("", productHandlerV1.PostProductV1)
			product.PUT("/:id", productHandlerV1.PutProductV1)
			product.DELETE("/:id", productHandlerV1.DeleteProductV1)
		}

		category := v1.Group("/categories")
		{
			categoryHandlerV1 := v1handler.NewCategoryHandler()
			category.GET("/:category", categoryHandlerV1.GetCategoryByCategoryV1)
		}

		news := v1.Group("/news")
		{
			newsHandlerV1 := v1handler.NewNewsHandler()
			// trick dùng parameter optional = DefaultQuery (Query optional)
			news.GET("", newsHandlerV1.GetNewsV1)
			news.GET("/:slug", newsHandlerV1.GetNewsV1)
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
