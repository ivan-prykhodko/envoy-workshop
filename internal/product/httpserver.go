package product

import "github.com/labstack/echo"

func NewHttpServer() *echo.Echo {
	srv := echo.New()
	srv.Debug = true

	pRouter := srv.Group("/products")
	pRouter.GET("/100500", getProductItem)

	return srv
}
