package product

import (
	"github.com/google/uuid"
	"github.com/jsn1096/ecommerce/model"
)

type UseCase interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Product, error)
	GetAll() (model.Products, error)
}

type Storage interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Product, error)
	GetAll() (model.Products, error)
}
