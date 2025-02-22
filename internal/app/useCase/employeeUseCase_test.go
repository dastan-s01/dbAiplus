package useCase

import (
	"context"
	"dbAiplus/internal/app/models"
	"dbAiplus/internal/app/repository"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

type mocks struct {
	employeeRepoMock repository.RepositoryMock
	sql              sqlmock.Sqlmock
}

func defaultMocks() *mocks {
	return &mocks{
		employeeRepoMock: repository.RepositoryMock{},
	}
}
func CreateEmployeeUseCase(mocks *mocks) EmployeeUseCase {
	return NewEmployeeUseCase(&mocks.employeeRepoMock)
}

var InvalidInput = errors.New("invalid input")
var InvalidResult = errors.New("invalid result")

func TestAddEmployeeName(t *testing.T) {
	mocks := defaultMocks()
	mocks.employeeRepoMock.CreateEmployeeFunc = func(ctx context.Context, employee models.Employee) error {
		if employee.Name != "TEST" || employee.Surname != "TEST" {
			return InvalidInput
		}
		return nil
	}
	uc := CreateEmployeeUseCase(mocks)
	err := uc.AddEmployee(context.Background(), models.Employee{
		Name:    "TEST",
		Surname: "TEST",
		City:    "TEST",
		Phone:   "+123456789",
	})
	if err != nil {
		t.Error(err)
		t.Error(InvalidResult)
	}
}
func TestAddEmployeeEmptyName(t *testing.T) {
	mocks := defaultMocks()
	mocks.employeeRepoMock.CreateEmployeeFunc = func(ctx context.Context, employee models.Employee) error {
		if employee.Name == "" || employee.Surname == "" || employee.City == "" {
			return InvalidInput
		}
		return nil
	}
	uc := CreateEmployeeUseCase(mocks)
	err := uc.AddEmployee(context.Background(), models.Employee{
		Name:    "",
		Surname: "",
		City:    "",
		Phone:   "+123456789",
	})
	if err == nil {
		t.Error("expected error but got nil")
	}
}
func TestAddEmployeePhone(t *testing.T) {
	mocks := defaultMocks()

	mocks.employeeRepoMock.CreateEmployeeFunc = func(ctx context.Context, employee models.Employee) error {
		if employee.Phone != "+123456789" {
			return InvalidInput
		}
		return nil
	}

	uc := CreateEmployeeUseCase(mocks)

	err := uc.AddEmployee(context.Background(), models.Employee{
		Name:    "TEST",
		Surname: "TEST",
		City:    "TEST",
		Phone:   "+123456789",
	})
	if err != nil {
		t.Error(err)
		t.Error(InvalidResult)
	}
}
func TestAddEmployeeInvalid(t *testing.T) {
	mocks := defaultMocks()

	mocks.employeeRepoMock.CreateEmployeeFunc = func(ctx context.Context, employee models.Employee) error {
		return errors.New("database error")
	}

	uc := CreateEmployeeUseCase(mocks)

	err := uc.AddEmployee(context.Background(), models.Employee{
		Name:    "TEST",
		Surname: "TEST",
		City:    "TEST",
		Phone:   "+123456789",
	})
	if err == nil {
		t.Error("expected error but got nil")
	}
}
