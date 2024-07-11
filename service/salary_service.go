package service

import (
	"context"
	"fmt"
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
	BulkAddSalaries(ctx context.Context, employee_salaries []entity.CreateEmployeeSalary) ([]entity.EmployeeSalary, error)
	GetAllSalary(ctx context.Context) ([]entity.EmployeeSalary, error)
	GetAllSalaryWithEmployee(ctx context.Context) ([]entity.EmployeeSalary, error)
	UpdateSalary(ctx context.Context, id string, employee_salary entity.CreateEmployeeSalary) (entity.EmployeeSalary, error)
	UpdateSalaryByEmployeeId(ctx context.Context, employee_salary entity.CreateEmployeeSalary) (entity.EmployeeSalary, error)
	DeleteSalary(ctx context.Context, id string, employee_salary entity.EmployeeSalary) error
	BulkDeleteSalaries(ctx context.Context, employee_salary_ids entity.BulkDeleteSalaries) error
}

func (s *SalaryService) AddSalaryService(ctx context.Context, employee_salary entity.CreateEmployeeSalary) (entity.EmployeeSalary, error) {
	//fmt.Println("FFFFFFFFFFFFFFFFFFFF")

	res, err := s.repository.AddSalary(ctx, employee_salary)
	if err != nil {
		log.Println("Error service function:", err)
	}
	//fmt.Println("WWWWWWWWWWWWW")
	return res, nil
}

func (s *SalaryService) BulkAddSalaryService(ctx context.Context, employee_salaries []entity.CreateEmployeeSalary) ([]entity.EmployeeSalary, error) {

	insertedSalaries, err := s.repository.BulkAddSalaries(ctx, employee_salaries)
	if err != nil {
		log.Panicln("Error service function:", err)
	}

	return insertedSalaries, nil
}

func (s *SalaryService) GetAllSalaryService(ctx context.Context) ([]entity.EmployeeSalary, error) {
	employee_salaries, err := s.repository.GetAllSalaryWithEmployee(ctx)
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

func (s *SalaryService) UpdateSalaryByIdService(ctx context.Context, employee_salary entity.CreateEmployeeSalary) (entity.EmployeeSalary, error) {

	//fmt.Println("Printing service employee salary:",employee_salary)

	res, err := s.repository.UpdateSalaryByEmployeeId(ctx, employee_salary)

	fmt.Println(res)

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

func (s *SalaryService) BulkDeleteSalaryService(ctx context.Context, employee_salary_ids entity.BulkDeleteSalaries) error {

	err := s.repository.BulkDeleteSalaries(ctx, employee_salary_ids)
	if err != nil {
		log.Println("Error service function:", err)
	}
	return nil
}
