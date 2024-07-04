package service

import (
	"context"
	"log"
	"salary_project/entity"
)

type SalaryService struct {
	repository Repository
}

func NewSalaryService(repository Repository) *SalaryService {
	return &SalaryService{
		repository: repository,
	}
}

type Repository interface {
	AddSalary(ctx context.Context, employee_salary entity.CreateEmployeeSalary) (entity.EmployeeSalary, error)
	GetAllSalary(ctx context.Context) ([]entity.EmployeeSalary, error)
	UpdateSalary(ctx context.Context, id string, employee_salary entity.CreateEmployeeSalary) (entity.EmployeeSalary, error)
	DeleteSalary(ctx context.Context, id string, employee_salary entity.EmployeeSalary) error
}

func (s *SalaryService) AddSalaryService(ctx context.Context, employee_salary entity.CreateEmployeeSalary) (entity.EmployeeSalary, error) {

	res, err := s.repository.AddSalary(ctx, employee_salary)
	if err != nil {
		log.Println("Error service function:", err)
	}
	return res, nil
}

func (s *SalaryService) GetAllSalaryService(ctx context.Context) ([]entity.EmployeeSalary, error) {
	employee_salaries, err := s.repository.GetAllSalary(ctx)
	if err != nil {
		log.Println("Error service function:", err)
	}

	return employee_salaries, nil
}

func (s *SalaryService) UpdateSalaryService(ctx context.Context, employee_salary entity.CreateEmployeeSalary, id string) (entity.EmployeeSalary, error) {

	res, err := s.repository.UpdateSalary(ctx, id, employee_salary)
	if err != nil {
		log.Println("Error service function:", err)
	}
	return res, nil
}

func (s *SalaryService) DeleteSalaryService(ctx context.Context, id string, employee_salary entity.EmployeeSalary) error {
	err := s.repository.DeleteSalary(ctx, id, employee_salary)
	if err != nil {
		log.Println("Error service function:", err)
	}
	return nil
}
