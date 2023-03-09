package handler

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jsn1096/ecommerce/infrastructure/handler/invoice"
	"github.com/jsn1096/ecommerce/infrastructure/handler/login"
	"github.com/jsn1096/ecommerce/infrastructure/handler/paypal"
	"github.com/jsn1096/ecommerce/infrastructure/handler/product"
	"github.com/jsn1096/ecommerce/infrastructure/handler/purchaseorder"
	"github.com/jsn1096/ecommerce/infrastructure/handler/user"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, dbPool *pgxpool.Pool) {
	health(e)

	invoice.NewRouter(e, dbPool)
	login.NewRouter(e, dbPool)
	paypal.NewRouter(e, dbPool)
	product.NewRouter(e, dbPool)
	purchaseorder.NewRouter(e, dbPool)
	user.NewRouter(e, dbPool)
}

func health(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			map[string]string{
				"time":         time.Now().String(),
				"message":      "hello World!",
				"service_name": "",
			},
		)
	})
}
