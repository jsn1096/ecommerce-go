package user

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jsn1096/ecommerce/model"
)

var (
	psqlInsert = "INSERT INTO users (id, email, password, details, created_at) VALUES ($1, $2, $3, $4, $5)"
	psqlGetAll = "SELECT id, email, password, details, created_at, updated_at FROM users"
)

// Pool de conexiones
// Esta estructura user a pesar de ser de postgres, ya implementa los
// métodos de la estructura storage xq tiene los mismos métodos, como
// las firmas coinciden con la interface, ya lo implementa
type User struct {
	db *pgxpool.Pool
}

// New returns a new User storage
func New(db *pgxpool.Pool) User {
	return User{db}
}

// Create a model.User
func (u User) Create(m *model.User) error {
	_, err := u.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Email,
		m.Password,
		m.IsAdmin,
		m.Details,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	query := psqlGetAll + " WHERE email = $1"
	row := u.db.QueryRow(
		context.Background(),
		query,
		email,
	)

	return u.scanRow(row)
}

// get all model.Users with fields
func (u User) GetAll() (model.Users, error) {
	rows, err := u.db.Query(
		context.Background(),
		psqlGetAll,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := model.Users{}
	for rows.Next() {
		m, err := u.scanRow(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return ms, nil
}

// Escanea las filas de la db y la convertimos a una estructura User
func (u User) scanRow(s pgx.Row) (model.User, error) {
	m := model.User{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Email,
		&m.Password,
		&m.IsAdmin,
		&m.Details,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}
	m.UpdatedAt = updatedAtNull.Int64

	return m, nil
}
