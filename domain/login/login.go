package login

import "github.com/jsn1096/ecommerce/model"

// crea el token
type UseCase interface {
	Login(email, password, jwtSecretKey string) (model.User, string, error)
}

// necesita el caso de uso de usuario
type UseCaseUser interface {
	Login(email, password string) (model.User, error)
}
