package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jsn1096/ecommerce/domain/user"
	storageUser "github.com/jsn1096/ecommerce/infrastructure/postgres/user"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	adminRoutes(e, h)
	publicRoutes(e, h)
}

// Contruimos el handler del caso de uso
func buildHandler(dbPool *pgxpool.Pool) handler {
	// Como se llaman igual el paquete user del domain y de infrastructure, lo llamamos storageUser al de infrast
	storage := storageUser.New(dbPool)
	useCase := user.New(storage)

	return newHandler(useCase)
}

// Construimos las rutas, publicas y administrativas
// En el siguiente módulo vamos a añadirle los permisos
func adminRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/admin/users")

	g.GET("", h.GetAll)
}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/users")

	g.POST("", h.Create)
}
