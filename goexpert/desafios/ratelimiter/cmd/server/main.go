package main

import (
	"github.com/osvaldotcf/pgfcycle/goexpert/desafio-ratelimiter/config"
	"github.com/osvaldotcf/pgfcycle/goexpert/desafio-ratelimiter/router"
)

func main() {
	config.Init()
	router.Init()
}
