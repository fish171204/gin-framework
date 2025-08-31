package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationError {
			switch e.Tag() {
			case "gt":
				errors[e.Field()] = fmt.Sprintf("%s phải lớn hơn giá trị tối thiểu là: %s", e.Field(), e.Param())
			case "lt":
				errors[e.Field()] = fmt.Sprintf("%s phải nhỏ hơn giá trị tối thiểu là: %s", e.Field(), e.Param())
			case "gte":
				errors[e.Field()] = fmt.Sprintf("%s phải lớn hơn hoặc bằng giá trị tối thiểu là: %s", e.Field(), e.Param())
			case "lte":
				errors[e.Field()] = fmt.Sprintf("%s phải lớn hơn hoặc bằng giá trị tối đa là: %s", e.Field(), e.Param())
			case "min":
				errors[e.Field()] = fmt.Sprintf("%s phải từ %s", e.Field(), e.Param())
			case "max":
				errors[e.Field()] = fmt.Sprintf("%s phải ít hơn %s", e.Field(), e.Param())
			// users
			case "uuid":
				errors[e.Field()] = e.Field() + " phải là UUID hợp lệ"
			// products
			case "slug":
				errors[e.Field()] = e.Field() + " chỉ được chứa chữ thường, số, dấu gạch ngang hoặc dấu chấm"
			case "required":
				errors[e.Field()] = e.Field() + " là bắt buộc"
			case "search":
				errors[e.Field()] = e.Field() + " chỉ được chứa chữ thường, in hoa ,số và khoảng trắng"
			case "email":
				errors[e.Field()] = e.Field() + " phải đúng định dạng email"
			case "datetime":
				errors[e.Field()] = e.Field() + " phải theo đúng định dạng YYYY-MM-DD"
			// category
			case "oneof":
				allowedValue := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[e.Field()] = fmt.Sprintf("%s phải là một trong các giá trị: %s", e.Field(), allowedValue)
			}

		}
		return gin.H{"error": errors}

	}

	return gin.H{"error": "Yêu cầu không hợp lệ " + err.Error()}
}

func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}

	// elon-musk-da- ....
	// [a-z0-9] = abc
	// [-.] = -
	// [-.][a-z0-9] = -abc
	// [a-z0-9]+(?:[-.][a-z0-9]+)*$ = abcas-abc-abac... (nhieu)
	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})

	var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		return searchRegex.MatchString(fl.Field().String())
	})
	return nil
}
