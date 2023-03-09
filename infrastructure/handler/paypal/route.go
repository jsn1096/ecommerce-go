package paypal

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jsn1096/ecommerce/domain/invoice"
	"github.com/jsn1096/ecommerce/domain/paypal"
	"github.com/jsn1096/ecommerce/domain/purchaseorder"
	storageInvoice "github.com/jsn1096/ecommerce/infrastructure/postgres/invoice"
	storagePurchaseOrder "github.com/jsn1096/ecommerce/infrastructure/postgres/purchaseorder"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	purchaseOrderUseCase := purchaseorder.New(storagePurchaseOrder.New(dbPool))
	invoiceUseCase := invoice.New(storageInvoice.New(dbPool), nil)
	useCase := paypal.New(purchaseOrderUseCase, invoiceUseCase)
	return newHandler(useCase)
}

// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(e *echo.Echo, h handler) {
	// Es recomendable que no se ponga paypal sino algo diferente como un uuid o algo para evitar posibles ataques
	route := e.Group("/api/v1/public/paypal")

	route.POST("", h.Webhook)
}
