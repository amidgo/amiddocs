package studydepstorage

import "github.com/amidgo/amiddocs/pkg/postgres"

type studyDepartmentStorage struct {
	p *postgres.Postgres
}

func New(p *postgres.Postgres) *studyDepartmentStorage {
	return &studyDepartmentStorage{p: p}
}
