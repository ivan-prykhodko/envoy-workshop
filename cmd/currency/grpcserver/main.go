package main

import "github.com/ivan-prykhodko/envoy-workshop/internal/currency"

func main() {
	ers := &currency.ExchangeRateService{}
	srv := currency.NewGrpcServer(9090, ers)
	srv.Run()
}
