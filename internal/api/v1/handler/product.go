package v1handler

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/fish171204/gin-framework/utils"
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
var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// PRODUCT V1
func (u *ProductHandler) GetProductV1(ctx *gin.Context) {
	search := ctx.Query("search")

	if err := utils.ValidationRequired("Search", search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(search) < 3 || len(search) > 50 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Search must be berween 3 and 50 characters"})
		return
	}

	if !searchRegex.MatchString(search) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Search must contain only letters, numbers and spaces"})
		return
	}

	// Đề bài: limit nếu trường hợp không tồn tại thì sẽ là 10 và phải là số nguyên dương
	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Limit must be a positive number"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products (V1)",
		"search":  search,
		"limit":   limit,
	})
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
