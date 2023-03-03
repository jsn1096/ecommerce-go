package purchaseorder

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jsn1096/ecommerce/domain/purchaseorder"
	"github.com/jsn1096/ecommerce/infrastructure/handler/response"
	"github.com/jsn1096/ecommerce/model"

	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase  purchaseorder.UseCase
	response response.API
}

func newHandler(useCase purchaseorder.UseCase) handler {
	return handler{useCase: useCase}
}

// Create handles the creation of a model.PurchaseOrder
func (h handler) Create(c echo.Context) error {
	m := model.PurchaseOrder{}
	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	userID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		return h.response.Error(c, "c.Get().(uuid.UUID)", errors.New("can´t parse uuid"))
	}

	m.UserID = userID
	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.response.Created(m))
}

// tarea: hacer un getAllPurchaseOrders para que el cliente pueda ver sus compras
