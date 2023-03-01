package user

import (
	"fmt"
	"net/http"

	"github.com/jsn1096/ecommerce/domain/user"
	"github.com/jsn1096/ecommerce/model"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase user.UseCase
}

func newHandler(uc user.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Create(c echo.Context) error {
	m := model.User{}

	// recibimos el objeto json que nos envia el cliente y lo convertimos a una estructura de go
	err := c.Bind(&m)
	if err != nil {
		fmt.Println(err)
	}
	// y le decimos al caso de uso que cree el usuario
	if err := h.useCase.Create(&m); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, m)
}

func (h handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}
