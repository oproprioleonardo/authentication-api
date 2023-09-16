package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/skyepic/privateapi/internal/infrastructure/dto/response"
)

var validate = validator.New()

func ValidateRequest[c any](str c) []*response.ValidatorErrorResponse {
	var errors []*response.ValidatorErrorResponse
	err := validate.Struct(str)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element response.ValidatorErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
