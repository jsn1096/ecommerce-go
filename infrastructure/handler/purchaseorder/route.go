package purchaseorder

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jsn1096/ecommerce/domain/purchaseorder"
	purchaseorderStorage "github.com/jsn1096/ecommerce/infrastructure/postgres/purchaseorder"

	"github.com/labstack/echo/v4"
)

// NewRouter returns a router to handle model.PurchaseOrder requests
func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	privateRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := purchaseorder.New(purchaseorderStorage.New(dbPool))
	return newHandler(useCase)
}

// privateRoutes handle the routes that requires a token
func privateRoutes(e *echo.Echo, h handler) {
	route := e.Group("/api/v1/private/purchase-orders")

	route.POST("", h.Create)
}
