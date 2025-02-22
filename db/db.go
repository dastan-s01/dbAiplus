package db

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectionDB(dsn string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	m, err := migrate.New("file://db/migrations", dsn)
	if err != nil {
		return nil, err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, err
	}
	
	return dbPool, nil
}
