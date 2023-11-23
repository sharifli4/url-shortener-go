package handler

import "github.com/gofiber/fiber/v2"

type HttpHandler struct {
}

func (h *HttpHandler) CreateURL(c *fiber.Ctx) error {
	c.Write([]byte("CreateURL"))
	return nil
}
