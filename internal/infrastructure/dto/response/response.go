package response

import (
	"github.com/gofiber/fiber/v2"
)

type (
	ValidatorErrorResponse struct {
		FailedField string
		Tag         string
		Value       string
	}
)

func NewResponseOKWithData(c *fiber.Ctx, data interface{}) error {
	response := fiber.Map{
		"error":   false,
		"status":  fiber.StatusOK,
		"data":    data,
		"message": "OK",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func NewResponseCreated(c *fiber.Ctx) error {
	response := fiber.Map{
		"error":   false,
		"status":  fiber.StatusCreated,
		"message": "Created",
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func NewResponseCreatedWithData(c *fiber.Ctx, data interface{}) error {
	response := fiber.Map{
		"error":   false,
		"status":  fiber.StatusCreated,
		"message": "Created",
		"data":    data,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func NewResponseCreatedWithURI(c *fiber.Ctx, uri string) error {
	response := fiber.Map{
		"error":   false,
		"status":  fiber.StatusCreated,
		"message": "Created",
		"uri":     uri,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func NewResponseUpdated(c *fiber.Ctx) error {
	response := fiber.Map{
		"error":   false,
		"status":  fiber.StatusOK,
		"message": "Updated",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func NewResponseUpdatedWithData(c *fiber.Ctx, data interface{}) error {
	response := fiber.Map{
		"data":    data,
		"error":   false,
		"status":  fiber.StatusOK,
		"message": "Updated",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func NewResponseOK(c *fiber.Ctx, message string) error {
	response := fiber.Map{
		"error":   false,
		"status":  fiber.StatusOK,
		"message": message,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func NewResponseBadRequest(c *fiber.Ctx, err interface{}) error {
	response := fiber.Map{
		"error":   true,
		"status":  fiber.StatusBadRequest,
		"message": err,
	}

	return c.Status(fiber.StatusBadRequest).JSON(response)
}

func NewResponseUnauthorized(c *fiber.Ctx, err interface{}) error {
	response := fiber.Map{
		"error":   true,
		"status":  fiber.StatusUnauthorized,
		"message": err,
	}

	return c.Status(fiber.StatusUnauthorized).JSON(response)
}

func NewResponseNoContent(c *fiber.Ctx) error {
	response := fiber.Map{
		"error":   false,
		"status":  fiber.StatusNoContent,
		"message": "No content",
	}

	return c.Status(fiber.StatusNoContent).JSON(response)
}

func NewResponseError(c *fiber.Ctx, err interface{}, code int) error {
	response := fiber.Map{
		"error":   true,
		"status":  code,
		"message": err,
	}

	return c.Status(code).JSON(response)
}

func NewValidatorErrorResponse(c *fiber.Ctx, tag string, failedField string, value string) error {
	response := ValidatorErrorResponse{
		FailedField: failedField,
		Tag:         tag,
		Value:       value,
	}

	return c.Status(fiber.StatusUnsupportedMediaType).JSON(response)
}
