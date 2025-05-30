package router

import (
	"github.com/go-chi/chi"
	"github.com/OsvaldoTCF/pgFCycle/goexpert/desafio-ratelimiter/limiter"
)

func Init() {
	router := chi.NewRouter()
	rate_limiter := limiter.InitializeRateLimiters()

	InitializeMiddlewares(router, rate_limiter)
	InitializeRoutes(router)
	InitilizeServer(router)
}
