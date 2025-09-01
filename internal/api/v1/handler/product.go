package v1handler

import (
	"net/http"
	"time"

	"github.com/fish171204/gin-framework/utils"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

type GetProductBySlugV1Param struct {
	Slug string `uri:"slug" binding:"slug,min=5,max=100"`
}

// omitempty = trống
// dive = kiểm tra phần bên trong (slice struct, map struct ,...)
type GetProductV1Param struct {
	Search string `form:"search" binding:"search,required,min=5,max=100"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
	Email  string `form:"email" binding:"omitempty,email"`
	Date   string `form:"date" binding:"omitempty,datetime=2006-01-02"`
}

type ProductImage struct {
	ImageName string `json:"image_name" binding:"required"`
	ImageLink string `json:"image_link" binding:"required,file_ext=jpg png gif"`
}

type ProductAttribute struct {
	AttributeName  string `json:"attribute_name" binding:"required"`
	AttributeValue string `json:"attribute_value" binding:"required"`
}

type ProductInfo struct {
	InfoKey   string `json:"info_key" binding:"required"`
	InfoValue string `json:"info_value" binding:"required"`
}
type PostProductV1Param struct {
	Name             string                 `json:"name" binding:"required,min=3,max=100"`
	Price            int                    `json:"price" binding:"required,min_int=100000"`
	Display          *bool                  `json:"display" binding:"omitempty"`
	ProductImage     ProductImage           `json:"product_image" binding:"required"`
	Tags             []string               `json:"tags" binding:"required,gt=3,lt=5"`
	ProductAttribute []ProductAttribute     `json:"product_attribute" binding:"required,gt=0,dive"`
	ProductInfo      map[string]ProductInfo `json:"product_info" binding:"required,gt=0,dive"`
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// PRODUCT V1
func (u *ProductHandler) GetProductV1(ctx *gin.Context) {
	var params GetProductV1Param
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	if params.Limit == 0 {
		params.Limit = 1
	}

	if params.Email == "" {
		params.Email = "No email"
	}

	if params.Date == "" {
		params.Date = time.Now().Format("2006-01-02")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products (V1)",
		"search":  params.Search,
		"limit":   params.Limit,
		"email":   params.Email,
		"date":    params.Date,
	})
}

func (u *ProductHandler) GetProductBySlugV1(ctx *gin.Context) {
	var param GetProductBySlugV1Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by Slug (V1)",
		"slug":    param.Slug,
	})
}

func (u *ProductHandler) PostProductV1(ctx *gin.Context) {
	var params PostProductV1Param
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	if params.Display == nil {
		defaultDisplay := true
		params.Display = &defaultDisplay
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":           "Create product (V1)",
		"name":              params.Name,
		"price":             params.Price,
		"display":           params.Display,
		"product_image":     params.ProductImage,
		"tags":              params.Tags,
		"product_attribute": params.ProductAttribute,
	})
}

func (u *ProductHandler) PutProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update product (V1)"})
}

func (u *ProductHandler) DeleteProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Delete product (V1)"})
}
