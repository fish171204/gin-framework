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
func (n *NewsHandler) PostNewsV1(ctx *gin.Context) {
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

	// Yêu cầu giới hạn file nhỏ hơn 5MB
	// 1 << 20 = 1 * 1^20 = 1 * 1048576 = 1MB
	// 1 << 20 = 5 * 1^20 = 5 * 1048576 = 5MB
	if image.Size > 5<<20 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File too large (5 MB)"})
		return
	}

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create upload folder"})
		return
	}

	dst := fmt.Sprintf("./uploads/%s", filepath.Base(image.Filename))

	if err := ctx.SaveUploadedFile(image, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot save file"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post news (V1)",
		"title":   params.Title,
		"status":  params.Status,
		"image":   image.Filename,
		"path":    dst,
	})
}

func (n *NewsHandler) PostUploadFileNewsV1(ctx *gin.Context) {
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

	filename, err := utils.ValidateAndSaveFile(image, "./uploads")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post news (V1)",
		"title":   params.Title,
		"status":  params.Status,
		"image":   filename,
		"path":    "./upload/" + filename,
	})
}

func (n *NewsHandler) PostUploadMultipleFileNewsV1(ctx *gin.Context) {
	var params PostNewsV1Param
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid multipart form"})
		return
	}

	// Lấy ra danh sách file từ form.File['images]
	images := form.File["images"]
	if len(images) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}

	var saveFile []string
	for _, image := range images {
		filename, err := utils.ValidateAndSaveFile(image, "./uploads")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		saveFile = append(saveFile, filename)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post news (V1)",
		"title":   params.Title,
		"status":  params.Status,
		"images":  saveFile,
	})
}
