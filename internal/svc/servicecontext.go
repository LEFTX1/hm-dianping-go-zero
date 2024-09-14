package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"hm-dianping-go-zero/internal/config"
	"hm-dianping-go-zero/internal/middleware"
)

type ServiceContext struct {
	Config                 config.Config
	RefreshTokenMiddleware rest.Middleware
	LoginMiddleware        rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                 c,
		RefreshTokenMiddleware: middleware.NewRefreshTokenMiddleware().Handle,
		LoginMiddleware:        middleware.NewLoginMiddleware().Handle,
	}
}
