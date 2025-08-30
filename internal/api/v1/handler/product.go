package v1handler

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

// elon-musk-da- ....
// [a-z0-9] = abc
// [-.] = -
// [-.][a-z0-9] = -abc
// [a-z0-9]+(?:[-.][a-z0-9]+)*$ = abcas-abc-abac... (nhieu)
var slugRegex = regexp.MustCompile(`[a-z0-9]+(?:[-.][a-z0-9]+)*$`)

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// PRODUCT V1
func (u *ProductHandler) GetProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List all products (V1)"})
}

func (u *ProductHandler) GetProductBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if !slugRegex.MatchString(slug) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Slug must contain only lowercase letter, number, hyphens and dots"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by Slug (V1)",
		"slug":    slug,
	})
}

func (u *ProductHandler) PostProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create product (V1)"})
}

func (u *ProductHandler) PutProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update product (V1)"})
}

func (u *ProductHandler) DeleteProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Delete product (V1)"})
}
