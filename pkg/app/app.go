package app

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/redis/go-redis/v9"
	"github.com/shari4ov/url-shortener-go/pkg/handler"
	"github.com/shari4ov/url-shortener-go/pkg/route"
	"github.com/shari4ov/url-shortener-go/pkg/service"
	"github.com/shari4ov/url-shortener-go/pkg/validation"
)

type App struct {
}

func (a *App) Start() {
	app := fiber.New()
	app.Use(logger.New())
	validate := validation.Validation{Validator: validator.New()}
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	urlService := service.UrlService{RedisService: redisClient, Ctx: context.Background()}
	handlers := handler.HttpHandler{Service: &urlService}
	routes := route.Route{
		App:           app,
		Handlers:      &handlers,
		UrlValidation: &validate,
	}
	routes.StartRoutes()
	app.Listen(":3000")
}
