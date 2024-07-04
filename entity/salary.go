package entity

import "time"

type EmployeeSalary struct {
	Id            int64     `json:"id"`
	Salary_Amount int       `json:"salary_amount"`
	Joining_Date  time.Time `json:"joining_date"`
	Project       string    `json:"project"`
}

type CreateEmployeeSalary struct {
	Salary_Amount int       `json:"salary_amount"`
	Joining_Date  time.Time `json:"joining_date"`
	Project       string    `json:"project"`
}
