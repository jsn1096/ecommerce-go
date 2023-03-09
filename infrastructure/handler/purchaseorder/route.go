package purchaseorder

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jsn1096/ecommerce/domain/purchaseorder"
	"github.com/jsn1096/ecommerce/infrastructure/handler/middle"
	purchaseorderStorage "github.com/jsn1096/ecommerce/infrastructure/postgres/purchaseorder"

	"github.com/labstack/echo/v4"
)

// NewRouter returns a router to handle model.PurchaseOrder requests
func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()

	privateRoutes(e, h, authMiddleware.IsValid)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := purchaseorder.New(purchaseorderStorage.New(dbPool))
	return newHandler(useCase)
}

// privateRoutes handle the routes that requires a token
func privateRoutes(e *echo.Echo, h handler, middleware ...echo.MiddlewareFunc) {
	route := e.Group("/api/v1/private/purchase-orders", middleware...)

	route.POST("", h.Create)
}
