package userstorage

import "github.com/amidgo/amiddocs/pkg/postgres"

type UserStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *UserStorage {
	return &UserStorage{p: p}
}
