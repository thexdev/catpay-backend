package main

import (
	"catpay/internal/infra/http"
)

func startHttpAdapter() {
	adapter := http.New()

	adapter.Bootstrap().Listen(":3000")
}

func main() {
	startHttpAdapter()
}
