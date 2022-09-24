package domain

import "github.com/jmoiron/sqlx"

type UserRepoDB struct {
	client *sqlx.DB
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(username string) (*User, error)
	Update(user User) (*User, error)
	Delete(user User) error
}
