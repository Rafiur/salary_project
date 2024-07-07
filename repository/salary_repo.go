package repository

import (
	"database/sql"
	"log"
	"salary_project/entity"
	"time"

	"golang.org/x/net/context"
)

type SalaryRepo struct {
	db *sql.DB
}

func NewSalaryRepo(db *sql.DB) *SalaryRepo {
	return &SalaryRepo{
		db: db,
	}
}

func (repo *SalaryRepo) AddSalary(ctx context.Context, employee_salary entity.CreateEmployeeSalary) (entity.EmployeeSalary, error) {
	var resposne entity.EmployeeSalary

	joining_date, err := time.Parse("2006-01-02", employee_salary.Joining_Date)
	if err != nil {
		log.Println("Error parsing joining date:", err)
	}

	qry := `INSERT INTO public.salary (salary_amount, joining_date, project) VALUES($1, $2, $3) RETURNING *`

	err = repo.db.QueryRowContext(ctx, qry, employee_salary.Salary_Amount, joining_date, employee_salary.Project).Scan(&resposne.Salary_Amount, &resposne.Joining_Date, &resposne.Project, &resposne.Id)

	return resposne, err
}

func (repo *SalaryRepo) GetAllSalary(ctx context.Context) ([]entity.EmployeeSalary, error) {
	qry := `SELECT id, salary_amount, joining_date, project FROM public.salary`

	rows, err := repo.db.QueryContext(ctx, qry)

	if err != nil {
		log.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	employee_salaries := []entity.EmployeeSalary{}

	for rows.Next() {
		var id int64
		var salary_amount int
		var joining_date string
		var project string

		err := rows.Scan(&id, &salary_amount, &joining_date, &project)
		if err != nil {
			log.Println("Error scanning row:", err)
		}
		employee_salary := entity.EmployeeSalary{Id: id, Salary_Amount: salary_amount, Joining_Date: joining_date, Project: project}
		employee_salaries = append(employee_salaries, employee_salary)
	}
	return employee_salaries, err
}

func (repo *SalaryRepo) UpdateSalary(ctx context.Context, id string, employee_salary entity.CreateEmployeeSalary) (entity.EmployeeSalary, error) {
	var resposne entity.EmployeeSalary

	qry := `UPDATE public.salary SET salary_amount=$1, joining_date=$2, project=$3 WHERE id=$4 RETURNING *`

	err := repo.db.QueryRowContext(ctx, qry, employee_salary.Salary_Amount, employee_salary.Joining_Date, employee_salary.Project, id).Scan(&resposne.Salary_Amount, &resposne.Joining_Date, &resposne.Project, &resposne.Id)

	if err != nil {
		log.Println("Error updating employee:", err)
	}
	return resposne, err
}

func (repo *SalaryRepo) DeleteSalary(ctx context.Context, id string, employee_salary entity.EmployeeSalary) error {
	qry := `DELETE from public.salary WHERE id=$1`

	_, err := repo.db.ExecContext(ctx, qry, id)
	if err != nil {
		log.Println("Error deleting employee:", err)
	}
	return err
}
