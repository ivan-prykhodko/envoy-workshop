package product

import (
	"context"
	"fmt"
	"math"

	"github.com/ivan-prykhodko/envoy-workshop/protogen/go/currency"
	"github.com/labstack/echo"
)

var grpcConn = NewGRpcConnection()

func getProductItem(c echo.Context) error {
	initialPrice := 123.45
	initialCurrency := "UAH"
	priceCurrency := "EUR"
	priceAmount, err := convert(initialPrice, initialCurrency, priceCurrency)
	if err != nil {
		fmt.Println("gRPC error:", err)
		priceAmount = initialPrice
		priceCurrency = initialCurrency
	}

	return c.JSON(200, map[string]any{
		"id":   "100500",
		"name": "Qwerty 100500",
		"price": map[string]any{
			"amount":   roundFloat(priceAmount, 2),
			"currency": priceCurrency,
		},
	})
}

func convert(amount float64, from, to string) (float64, error) {
	rate, err := exchangeRate(context.Background(), from, to)
	if err != nil {
		return 0, err
	}
	return amount * rate, nil
}

func exchangeRate(ctx context.Context, from, to string) (float64, error) {
	clientConn, err := grpcConn.ClientConn()
	if err != nil {
		return 0, err
	}

	svc := currency.NewExchangeRateServiceClient(clientConn)
	req := &currency.ExchangeRateRequest{
		FromCurrency: from,
		ToCurrency:   to,
	}
	resp, err := svc.GetExchangeRate(ctx, req)
	if err != nil {
		return 0, err
	}

	return resp.ExchangeRate.Rate, nil
}

func roundFloat(val float64, scale uint) float64 {
	ratio := math.Pow(10, float64(scale))
	return math.Round(val*ratio) / ratio
}
