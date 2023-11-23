package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shari4ov/url-shortener-go/pkg/httpkit"
)

type HttpHandler struct {
}

func (h *HttpHandler) CreateURL(c *fiber.Ctx) error {
	var p httpkit.UrlPayload
	if err := c.BodyParser(&p); err != nil {
		return err
	}
	responsePayload := httpkit.ResponsePayload{
		ShortUrl: "http://localhost:3000/123456",
		Status:   fiber.StatusOK,
	}
	return c.Status(fiber.StatusOK).JSON(responsePayload)
}
