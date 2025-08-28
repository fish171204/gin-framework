package lesson

import (
	"github.com/gin-gonic/gin"
)

func Lesson02GinStarter() {
	r := gin.Default()
	r.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Nguyen Dang Khoa"})
	})

	r.Run(":8080")
}
