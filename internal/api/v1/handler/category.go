package v1handler

import (
	"net/http"

	"github.com/fish171204/gin-framework/utils"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

type GetCategoryByCategoryV1Param struct {
	Category string `uri:"category" binding:"oneof=php python golang"`
}

type PostCategoryV1Param struct {
	Name   string `form:"name" binding:"required"`
	Status string `form:"status" binding:"required,oneof=1 2"`
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (u *CategoryHandler) GetCategoryByCategoryV1(ctx *gin.Context) {
	var params GetCategoryByCategoryV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Category found",
		"category": params.Category,
	})
}

// func (u *CategoryHandler) PostCategoryV1(ctx *gin.Context) {
// 	var params PostCategoryV1Param
// 	if err := ctx.ShouldBind(&params); err != nil {
// 		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "Post category (V1)",
// 		"name":    params.Name,
// 		"status":  params.Status,
// 	})
// }
