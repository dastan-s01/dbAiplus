package useCase

import (
	"context"
	"dbAiplus/internal/app/models"
	"dbAiplus/internal/app/repository"
)

type EmployeeUseCase interface {
	AddEmployee(ctx context.Context, employee models.Employee) error
}

type employeeUseCase struct {
	repo repository.Repository
}

func NewEmployeeUseCase(repo repository.Repository) EmployeeUseCase {
	return &employeeUseCase{repo: repo}
}

func (uc *employeeUseCase) AddEmployee(ctx context.Context, employee models.Employee) error {
	return uc.repo.CreateEmployee(ctx, employee)
}
