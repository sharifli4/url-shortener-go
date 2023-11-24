package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shari4ov/url-shortener-go/pkg/handler"
	"github.com/shari4ov/url-shortener-go/pkg/validation"
)

type Route struct {
	App           *fiber.App
	Handlers      *handler.HttpHandler
	UrlValidation *validation.Validation
}

func (r *Route) StartRoutes() {
	r.App.Post("/", r.UrlValidation.ValidateURLPayload, r.Handlers.CreateURL)
	r.App.Get("/:token", r.Handlers.RedirectURL)
}
