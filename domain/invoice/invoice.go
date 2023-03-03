package invoice

import (
	"github.com/jsn1096/ecommerce/model"
)

type UseCase interface {
	Create(m *model.PurchaseOrder) error
}

type Storage interface {
	Create(m *model.Invoice, ms model.InvoiceDetails) error
}
