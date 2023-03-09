package paypal

import (
	"io"
	"log"
	"net/http"

	"github.com/jsn1096/ecommerce/domain/paypal"
	"github.com/jsn1096/ecommerce/infrastructure/handler/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCasePayPal paypal.UseCase
	responser     response.API
}

func newHandler(ucp paypal.UseCase) handler {
	return handler{useCasePayPal: ucp}
}

// recibe el body de la petición que hacen de paypal hacia nosotros
// lo convertimos en un slice de bytes
func (h handler) Webhook(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		h.responser.BindFailed(err)
	}
	// por medio de go rutinas proceso el webhook, que es que realiza
	// la función asíncrona, ya que no me interesa que paypal espere que
	// procese el pago, sino solo saber que me llegó el webhook, ya
	// si lo puedo procesar o no es nuestro problema
	go func() {
		err = h.useCasePayPal.ProcessRequest(c.Request().Header, body)
		if err != nil {
			log.Printf("useCasePayPal.ProccesRequest(): %v", err)
		}
	}()
	// respondemos
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
