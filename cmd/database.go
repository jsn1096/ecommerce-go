package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

const AppName = "EDcommerce"

func newDBConection() (*pgxpool.Pool, error) {
	// Numero de conexiones disponibles
	min := 3
	max := 100

	minConn := os.Getenv("DB_MIN_CONN")
	maxConn := os.Getenv("DB_MAX_CONN")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	// validamos que si se quiere modificar el numero de conexiones, nunca pasen del minimo ni máximo
	if minConn != "" {
		v, err := strconv.Atoi(minConn)
		if err != nil {
			log.Println("Warning: DB_MIN_CONN has not a valid value, we will set min connections to", min)
		} else {
			if v >= min && v <= max {
				min = v
			}
		}
	}
	if maxConn != "" {
		v, err := strconv.Atoi(maxConn)
		if err != nil {
			log.Println("Warning: DB_MAX_CONN has not a valid value, we will set max connections to", max)
		} else {
			if v >= min && v <= max {
				max = v
			}
		}
	}

	// Obtenemos la url para conectarnos a la db
	connString := makeURL(user, pass, host, port, dbName, sslMode, min, max)
	// Si la conexión a la db requiere sslmode, tenemos que darle a esta url
	// la ubicación del archivo de conexión
	if os.Getenv("DB_SSL_MODE") == "require" {
		connString += " sslrootcert=" + os.Getenv("DB_SSL_ROOT_CERT")
	}
	// Valida que la url esté correctamente formada
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("%s %w", "pgxpool.ParseConfig()", err)
	}

	// cuando se haga el trace a una db se ve el nombre que le pusimos
	config.ConnConfig.RuntimeParams["application_name"] = AppName

	// Creamos el pool de conexiones, le enviamos un contexto vacío y la configuracion que creamos
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("%s %w", "pgxpool.NewWithConfig()", err)
	}
	// si todo está bien devolvemos el pool de conexiones
	return pool, nil
}

// Creamos la url que tenemos que pasarle al pool de conexiones
func makeURL(user, pass, host, port, dbName, sslMode string, minConn, maxConn int) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_min_conns=%d pool_max_conns=%d",
		user,
		pass,
		host,
		port,
		dbName,
		sslMode,
		minConn,
		maxConn,
	)
}
