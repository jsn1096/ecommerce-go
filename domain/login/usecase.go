package login

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jsn1096/ecommerce/model"
)

type Login struct {
	useCaseUser UseCaseUser
}

func New(uc UseCaseUser) Login {
	return Login{useCaseUser: uc}
}

func (l Login) Login(email, password, jwtSecretKey string) (model.User, string, error) {
	// recibimos el usuario del caso de uso que este ya valida si existe y todo eso
	user, err := l.useCaseUser.Login(email, password)
	if err != nil {
		return model.User{}, "", fmt.Errorf("%s %w", "useCaseUser.Login()", err)
	}
	// llenamos el modelo jwtcustomclaims con los datos del usuario recibido
	claims := model.JWTCustomClaims{
		UserID:  user.ID,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	// Creamos el token con los datos del claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// firmamos el token con nuestra variable de etorno secret-key
	data, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return model.User{}, "", fmt.Errorf("%s %w", "token.SignedString()", err)
	}

	user.Password = ""
	// devolvemos el usuario y el token firmado, y el error nulo
	return user, data, nil
}
