package v1handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fish171204/gin-framework/utils"
	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
}

type PostNewsV1Param struct {
	Title  string `form:"title" binding:"required"`
	Status string `form:"status" binding:"required,oneof=1 2"`
}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

func (u *NewsHandler) GetNewsV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if slug == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get News (V1)",
			"slug":    "No News",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get News (V1)",
			"slug":    slug,
		})
	}
}

// // Body -> form-data
func (u *NewsHandler) PostNewsV1(ctx *gin.Context) {
	var params PostNewsV1Param
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// Trường hợp chưa có folder -> Tạo folder
	// os.ModePerm = 0777 (octal)
	// Có nghĩa: đọc, ghi, thực thi (read, write, execute) cho tất cả mọi người (owner, group, others)
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create upload folder"})
	}

	dst := fmt.Sprintf("./uploads/%s", filepath.Base(image.Filename))

	if err := ctx.SaveUploadedFile(image, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot save file"})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post category (V1)",
		"title":   params.Title,
		"status":  params.Status,
		"image":   image.Filename,
	})

}
