package di

import (
	"dbAiplus/internal/app/repository"
	"dbAiplus/internal/app/useCase"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DI struct {
	EmployeeUseCase useCase.EmployeeUseCase
}

func NewDI(db *pgxpool.Pool) *DI {
	repo := repository.NewRepository(db)
	useCase := useCase.NewEmployeeUseCase(repo)

	return &DI{
		EmployeeUseCase: useCase,
	}
}
