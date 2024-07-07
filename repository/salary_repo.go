package repository

import (
	"database/sql"
	"fmt"
	"log"
	"salary_project/entity"

	"github.com/lib/pq"
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

	// joining_date, err := time.Parse("2006-01-02", employee_salary.Joining_Date)
	// if err != nil {
	// 	log.Println("Error parsing joining date:", err)
	// }

	qry := `INSERT INTO public.salary (salary_amount, joining_date, project) VALUES($1, $2, $3) RETURNING *`

	err := repo.db.QueryRowContext(ctx, qry, employee_salary.Salary_Amount, employee_salary.Joining_Date, employee_salary.Project).
		Scan(&resposne.Salary_Amount, &resposne.Joining_Date, &resposne.Project, &resposne.Id)

	return resposne, err
}

func (repo *SalaryRepo) BulkAddSalaries(ctx context.Context, employee_salaries []entity.CreateEmployeeSalary) ([]entity.EmployeeSalary, error) {

	// Begin a transaction
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error executing bulk add transaction:", err)
	}
	defer func() {
		if err != nil {
			fmt.Println("rolling back")
			tx.Rollback()
			return
		}
		//fmt.Println("Commiting")
		err = tx.Commit()
	}()
	//fmt.Println(employee_salaries)
	// Prepare the SQL statement for bulk insert
	qry := `INSERT INTO public.salary (salary_amount, joining_date, project) VALUES ($1, $2, $3) RETURNING id, salary_amount, joining_date, project`
	stmt, err := tx.PrepareContext(ctx, qry)
	if err != nil {
		log.Println("Error querying database:", err)
	}
	defer stmt.Close()

	var inserted_salaries []entity.EmployeeSalary

	// Iterate over each salary and execute the insert statement
	for _, employee_salary := range employee_salaries {
		var inserted_salary entity.EmployeeSalary
		//fmt.Println("inside insert for loop")
		err := stmt.QueryRowContext(ctx, employee_salary.Salary_Amount, employee_salary.Joining_Date, employee_salary.Project).
			Scan(&inserted_salary.Id, &inserted_salary.Salary_Amount, &inserted_salary.Joining_Date, &inserted_salary.Project)
		if err != nil {
			log.Println("Error iterating over salaries:", err)
			continue
		}
		//fmt.Println("inserted salary: ", inserted_salary)
		inserted_salaries = append(inserted_salaries, inserted_salary)
	}
	//fmt.Println("inserted salaries", inserted_salaries)
	return inserted_salaries, nil
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
		log.Println("Error updating salary:", err)
	}
	return resposne, err
}

func (repo *SalaryRepo) DeleteSalary(ctx context.Context, id string, employee_salary entity.EmployeeSalary) error {
	qry := `DELETE from public.salary WHERE id=$1`

	_, err := repo.db.ExecContext(ctx, qry, id)
	if err != nil {
		log.Println("Error deleting salary:", err)
	}
	return err
}

func (repo *SalaryRepo) BulkDeleteSalaries(ctx context.Context, employee_salary_ids entity.BulkDeleteSalaries) error {

	// Begin a transaction
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error executing bulk delete transaction:", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Prepare the SQL statement for bulk delete
	qry := `DELETE FROM public.salary WHERE id = ANY($1)`
	stmt, err := tx.PrepareContext(ctx, qry)
	if err != nil {
		log.Println("Error quering database:", err)
	}
	defer stmt.Close()

	// need to perform conversion for pq.Array
	// Convert []int64 to []interface{} for pq.Array
    ids := make([]interface{}, len(employee_salary_ids.Ids))
    for i, id := range employee_salary_ids.Ids {
        ids[i] = id
    }

    // Execute the delete statement with pq.Array for array literal
    _, err = stmt.ExecContext(ctx, pq.Array(ids))
    if err != nil {
        log.Println("Error executing delete statement:", err)
        return err
    }

    return nil
}
