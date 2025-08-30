package v1handler

import (
	"net/http"

	"github.com/fish171204/gin-framework/utils"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

var validCategory = map[string]bool{
	"php":    true,
	"golang": true,
	"python": true,
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (u *CategoryHandler) GetCategoryByCategoryV1(ctx *gin.Context) {
	category := ctx.Param("category")

	if err := utils.ValidationInList("Category", category, validCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Category found",
		"category": category,
	})
}
