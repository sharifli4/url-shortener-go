package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/shari4ov/url-shortener-go/pkg/config"
	"github.com/shari4ov/url-shortener-go/pkg/httpkit"
	"github.com/shari4ov/url-shortener-go/pkg/service"
	"github.com/shari4ov/url-shortener-go/pkg/token"
)

type HttpHandler struct {
	Service *service.UrlService
}

func (h *HttpHandler) CreateURL(c *fiber.Ctx) error {
	var p httpkit.UrlPayload
	if err := c.BodyParser(&p); err != nil {
		return err
	}

	tokenUrl := service.UrlToken{
		Url:   p.Url,
		Token: token.GenerateToken(),
	}
	h.Service.SaveToken(tokenUrl)
	responsePayload := httpkit.ResponsePayload{
		ShortUrl: fmt.Sprintf("%s/%s", config.TinyDomain, tokenUrl.Token),
		Status:   fiber.StatusOK,
	}
	return c.Status(fiber.StatusOK).JSON(responsePayload)
}
func (h *HttpHandler) RedirectURL(c *fiber.Ctx) error {
	userToken := c.Params("token")
	url := h.Service.GetUrl(userToken)
	if url == "" {
		errorPayload := httpkit.ErrorPayload{
			Message: "Url not found",
			Status:  fiber.StatusNotFound,
		}
		return c.Status(fiber.StatusNotFound).JSON(errorPayload)
	}
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}
