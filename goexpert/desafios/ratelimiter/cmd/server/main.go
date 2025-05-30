package main

import (
	"github.com/OsvaldoTCF/pgFCycle/goexpert/desafio-ratelimiter/config"
	"github.com/OsvaldoTCF/pgFCycle/goexpert/desafio-ratelimiter/router"
)

func main() {
	config.Init()
	router.Init()
}
