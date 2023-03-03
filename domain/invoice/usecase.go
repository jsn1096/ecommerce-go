package invoice

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jsn1096/ecommerce/model"
)

// Invoice implements UseCase
type Invoice struct {
	storage Storage
}

// New returns a new Invoice
func New(s Storage) Invoice {
	return Invoice{storage: s}
}

// Create creates a model.Invoice
func (i Invoice) Create(po *model.PurchaseOrder) error {
	// Valida si tiene la estructura correcta
	if err := po.Validate(); err != nil {
		return fmt.Errorf("invoice: %w", err)
	}
	// Creamos la factura y los detalles a partir de la orden de compra que recibimos
	invoice, invoiceDetails, err := invoiceFromPurchaseOrder(po)
	if err != nil {
		return fmt.Errorf("%s %w", "invoiceFromPurchaseOrder()", err)
	}
	// Aquí ya le decimos al storage que cree la factura y detalles
	err = i.storage.Create(&invoice, invoiceDetails)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	return nil
}

func invoiceFromPurchaseOrder(po *model.PurchaseOrder) (model.Invoice, model.InvoiceDetails, error) {
	ID, err := uuid.NewUUID()
	if err != nil {
		return model.Invoice{}, nil, fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}
	// Encabezado de la factura
	invoice := model.Invoice{
		ID:              ID,
		UserID:          po.UserID,
		PurchaseOrderID: po.ID,
		CreatedAt:       time.Now().Unix(),
	}

	// le agignamos a products los productos de la orden de compra
	var products model.ProductToPurchases
	// transformamos de json a una estructura de go
	err = json.Unmarshal(po.Products, &products)
	if err != nil {
		return model.Invoice{}, nil, fmt.Errorf("%s %w", "json.Unmarshal()", err)
	}

	// Detalle de la factura
	var invoiceDetails model.InvoiceDetails
	for _, v := range products {
		detailID, err := uuid.NewUUID()
		if err != nil {
			return model.Invoice{}, nil, fmt.Errorf("%s %w", "uuid.NewUUID()", err)
		}

		detail := model.InvoiceDetail{
			ID:        detailID,
			InvoiceID: invoice.ID,
			ProductID: v.ProductID,
			Amount:    v.Amount,
			UnitPrice: v.UnitPrice,
			CreatedAt: time.Now().Unix(),
		}

		invoiceDetails = append(invoiceDetails, detail)
	}

	return invoice, invoiceDetails, nil
}
