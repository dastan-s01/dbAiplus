package repository

import (
	"context"
	"dbAiplus/internal/app/models"
	"github.com/jackc/pgconn"
)

//go:generate moq -out repo_mock.go . Repository
type DB interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}
type Repository interface {
	CreateEmployee(ctx context.Context, employee models.Employee) error
}
type repository struct {
	db DB
}

func NewRepository(db DB) *repository {
	return &repository{
		db: db,
	}
}
func (r *repository) CreateEmployee(ctx context.Context, employee models.Employee) error {
	query := `INSERT INTO employees (Name,Surname,City,Phone) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, employee.Name, employee.Surname, employee.City, employee.Phone)
	return err
}
