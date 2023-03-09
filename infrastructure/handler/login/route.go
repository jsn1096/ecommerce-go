package login

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jsn1096/ecommerce/domain/login"
	"github.com/jsn1096/ecommerce/domain/user"

	userStorage "github.com/jsn1096/ecommerce/infrastructure/postgres/user"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	// build middlewares to validate permissions on the roecho

	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCaseUser := user.New(userStorage.New(dbPool))
	useCase := login.New(useCaseUser)
	return newHandler(useCase)
}

// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/api/v1/public/login")

	route.POST("", h.Login)
}
