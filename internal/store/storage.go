package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Users interface {
		Create(context.Context) error
	}
	Posts interface {
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UsersStore{db},
		Posts: &PostsStore{db},
	}
}
