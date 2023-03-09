package main

import (
	"log"
	"os"

	"github.com/jsn1096/ecommerce/infrastructure/handler"
	"github.com/jsn1096/ecommerce/infrastructure/handler/response"
)

func main() {
	// Si no cargan nuestras variables de entorno salimos
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}
	// If the environment variables have not values, break
	err = validateEnvironments()
	if err != nil {
		log.Fatal(err)
	}
	// Create Echo for routing
	e := newHTTP(response.HTTPErrorHandler)

	dbPool, err := newDBConection()
	if err != nil {
		log.Fatal(err)
	}

	handler.InitRoutes(e, dbPool)

	port := os.Getenv("SERVER_PORT")
	if os.Getenv("IS_HTTPS") == "true" {
		err = e.StartTLS(":"+port, os.Getenv("CERT_PEM_FILE"), os.Getenv("KEY_PEM"))
	} else {
		err = e.Start(":" + port)
	}
	if err != nil {
		log.Fatal(err)
	}
}
