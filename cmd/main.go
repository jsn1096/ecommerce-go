package main

import (
	"log"
	"os"

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

	_ = dbPool

	err = e.Start(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
