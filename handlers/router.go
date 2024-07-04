package handlers

import (
	"database/sql"
	"salary_project/repository"
	"salary_project/service"

	"github.com/labstack/echo/v4"
)

func SetUpRouter(e *echo.Echo, db *sql.DB) {

	salaryRepo := repository.NewSalaryRepo(db)
	salaryService := service.NewSalaryService(salaryRepo)
	salaryHandler := NewSalaryHandler(salaryService)


	salaryGroup := e.Group("/employee/salary")

	salaryHandler.MapSalaryRoutes(salaryGroup)

}