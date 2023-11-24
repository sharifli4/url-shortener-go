package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/shari4ov/url-shortener-go/pkg/httpkit"
)

type Validation struct {
	Validator *validator.Validate
}

func (v *Validation) ValidateURLPayload(c *fiber.Ctx) error {
	body := new(httpkit.UrlPayload)
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	err := v.Validator.Struct(body)
	if err != nil {
		payloadError := httpkit.ErrorPayload{
			Message: "Invalid Payload",
			Status:  fiber.StatusBadRequest,
		}
		return c.Status(fiber.StatusBadRequest).JSON(payloadError)
	}
	return c.Next()
}
