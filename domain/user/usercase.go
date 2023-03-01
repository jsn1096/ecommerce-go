package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jsn1096/ecommerce/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{storage: s}
}

func (u User) Create(m *model.User) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}
	m.ID = ID
	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword()", err)
	}

	// el password que nos devuelve es un slice de bytes x lo que se debe convertir a string
	m.Password = string(password)
	if m.Details == nil {
		m.Details = []byte("{}")
	}

	m.CreatedAt = time.Now().Unix()

	// Aquí lo guardamos en la base de datos
	err = u.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.create()", err)
	}

	m.Password = ""
	return nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	user, err := u.storage.GetByEmail(email)
	if err != nil {
		return model.User{}, err //fmt.Errorf("#{"storage.GetByEmail()"} #{err}")
	}
	return user, nil
}

func (u User) GetAll() (model.Users, error) {
	users, err := u.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "storage.GetAll()", err)
	}

	return users, nil
}
