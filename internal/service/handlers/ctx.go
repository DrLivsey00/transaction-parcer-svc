package handlers

import (
	"context"
	"net/http"

	"github.com/DrLivsey00/transaction-parcer-svc/internal/config"
	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/services"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int
type serviceKey int
type configKey int

const (
	logCtxKey     ctxKey     = iota
	serviceCtxKey serviceKey = iota
	configCtxKey  configKey  = iota
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxService(service *services.Services) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, serviceCtxKey, service)
	}
}

func Service(r *http.Request) *services.Services {
	return r.Context().Value(serviceCtxKey).(*services.Services)
}

func CtxConfig(cfg config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, configCtxKey, cfg)
	}
}

func GetConfig(r *http.Request) config.Config {
	return r.Context().Value(configCtxKey).(config.Config)
}
