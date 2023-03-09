package login

import (
	"database/sql"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/jsn1096/ecommerce/domain/login"
	"github.com/jsn1096/ecommerce/infrastructure/handler/response"
	"github.com/jsn1096/ecommerce/model"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase   login.UseCase
	responser response.API
}

func newHandler(useCase login.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Login(c echo.Context) error {
	m := model.Login{}
	// recibimos la petición del cliente
	err := c.Bind(&m)
	if err != nil {
		return h.responser.BindFailed(err)
	}
	// usamos el login del caso de uso, cogemos el handler de las variables de entorno
	u, t, err := h.useCase.Login(m.Email, m.Password, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		// si no coincide el usuario o la contraseña lo correcto es mandar el mensaje que estamos mandando
		if strings.Contains(err.Error(), "bcrypt.CompareHashAndPassword()") ||
			errors.Is(err, sql.ErrNoRows) {
			resp := model.MessageResponse{
				Data:   "wrong user or password",
				Errors: model.Responses{{Code: response.AuthError, Message: "wrong user or password"}},
			}
			return c.JSON(http.StatusBadRequest, resp)
		}
		return h.responser.Error(c, "useCase.Login()", err)
	}
	// Si todo está bien devolvemos el usuario y el token
	return c.JSON(h.responser.OK(map[string]interface{}{"user": u, "token": t}))
}
