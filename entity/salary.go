package entity

type EmployeeSalary struct {
	Id            int64  `json:"id"`
	Salary_Amount int    `json:"salary_amount"`
	Joining_Date  string `json:"joining_date"`
	Project       string `json:"project"`
	Employee_Id   int32  `json:"employee_id"`
}

type CreateEmployeeSalary struct {
	Salary_Amount int    `json:"salary_amount"`
	Joining_Date  string `json:"joining_date"`
	Project       string `json:"project"`
	Employee_Id   int32  `json:"employee_id"`
}

type BulkCreateEmployeeSalary struct {
	BulkSalaries []CreateEmployeeSalary `json:"employee_salaries"`
}

type BulkDeleteSalaries struct {
	Ids []int64 `json:"ids"`
}
