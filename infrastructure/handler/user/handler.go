package user

import (
	"github.com/jsn1096/ecommerce/domain/user"
	"github.com/jsn1096/ecommerce/infrastructure/handler/response"
	"github.com/jsn1096/ecommerce/model"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase   user.UseCase
	responser response.API
}

func newHandler(uc user.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Create(c echo.Context) error {
	m := model.User{}

	// recibimos el objeto json que nos envia el cliente y lo convertimos a una estructura de go
	err := c.Bind(&m)
	if err != nil {
		return h.responser.BindFailed(err)
	}
	// y le decimos al caso de uso que cree el usuario
	if err := h.useCase.Create(&m); err != nil {
		return h.responser.Error(c, "useCase.Create()", err)
	}
	return c.JSON(h.responser.Created(m))
}

func (h handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return h.responser.Error(c, "useCase.GetAll()", err)
	}
	return c.JSON(h.responser.OK(users))
}
