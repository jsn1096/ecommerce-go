package middle

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/jsn1096/ecommerce/infrastructure/handler/response"
	"github.com/jsn1096/ecommerce/model"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	responser response.API
}

func New() AuthMiddleware {
	return AuthMiddleware{}
}

// Middleware que se ejecuta antes de la funcion principal
func (am AuthMiddleware) IsValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obtenemos el token del request
		token, err := getTokenFromRequest(c.Request())
		if err != nil {
			return am.responser.BindFailed(err)
		}
		// validamos el token
		isValid, claims := am.validate(token)
		if !isValid {
			err = errors.New("the token is not valid")
			return am.responser.BindFailed(err)
		}
		// si es válido le agregamos los claims al contexto
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("isAdmin", claims.IsAdmin)

		return next(c)
	}
}

// Valida si es administrador o no
func (am AuthMiddleware) IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin, ok := c.Get("isAdmin").(bool)
		if !isAdmin || !ok {
			err := errors.New("you are not admin")
			return am.responser.BindFailed(err)
		}

		return next(c)
	}
}

// Valida si el token es válido
func (am AuthMiddleware) validate(token string) (bool, model.JWTCustomClaims) {
	// Esta funcion propia de jwt es la que valida el token
	claims, err := jwt.ParseWithClaims(token, &model.JWTCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		log.Println(token)
		log.Println(os.Getenv("JWT_SECRET_KEY"))
		log.Println(err)
		return false, model.JWTCustomClaims{}
	}
	// le asignamos a data los customclaims que recibimos
	data, ok := claims.Claims.(*model.JWTCustomClaims)
	if !ok {
		log.Println("is not a jwtcustomclaims")
		return false, model.JWTCustomClaims{}
	}
	// devolvemos que sí es válido y los customclaims
	return true, *data
}

// obtiene el token del request
func getTokenFromRequest(r *http.Request) (string, error) {
	// buscamos el header de authorization
	data := r.Header.Get("Authorization")
	if data == "" {
		return "", errors.New("el header de autorización está vacío")
	}
	// si viene la palabra Bearer token tambien es válido, esta funcion es para
	// quitar esa palabra
	if strings.HasPrefix(data, "Bearer") {
		return data[7:], nil
	}

	return data, nil
}
