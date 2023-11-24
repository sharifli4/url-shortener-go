package service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type UrlService struct {
	RedisService *redis.Client
	Ctx          context.Context
}
type UrlToken struct {
	Token string
	Url   string
}

func (us *UrlService) SaveToken(uT UrlToken) {
	us.RedisService.Set(us.Ctx, uT.Token, uT.Url, 15*time.Minute)
}

func (us *UrlService) GetUrl(token string) string {
	return us.RedisService.Get(us.Ctx, token).Val()
}
