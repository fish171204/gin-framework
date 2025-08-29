package lesson

import (
	"github.com/gin-gonic/gin"
)

func Lesson02GinStarter() {
	r := gin.Default()
	r.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Nguyen Dang Khoa"})
	})

	r.GET("/users/:user_id", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "Danh sach user"})
	})

	r.GET("/user/:user_id", func(ctx *gin.Context) {
		user_id := ctx.Param("user_id")
		ctx.JSON(200, gin.H{
			"data":    "Thong tin user",
			"user_id": user_id,
		})
	})

	r.GET("/products/detail/:product_name", func(ctx *gin.Context) {
		product_name := ctx.Param("product_name")
		page := ctx.Query("page")
		ctx.JSON(200, gin.H{
			"data:":        "Thong tin san pham",
			"product_name": product_name,
			"page":         page,
		})
	})

	r.GET("/product/:product_name", func(ctx *gin.Context) {
		product_name := ctx.Param("product_name")
		price := ctx.Query("price")
		color := ctx.Query("color")
		ctx.JSON(200, gin.H{
			"data":          "Thong tin san pham",
			"product_name":  product_name,
			"product_price": price,
			"color":         color,
		})
	})
	r.Run(":8080")
}
