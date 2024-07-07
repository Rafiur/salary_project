package handlers

import (
	"fmt"
	"log"
	"net/http"
	"salary_project/entity"
	"salary_project/service"

	"github.com/labstack/echo/v4"
)

type SalaryHandler struct {
	SalaryService service.SalaryService
}

func NewSalaryHandler(salaryService *service.SalaryService) *SalaryHandler {
	return &SalaryHandler{
		SalaryService: *salaryService,
	}
}

func (h *SalaryHandler) MapSalaryRoutes(salaryGroup *echo.Group) {

	salaryGroup.POST("/add", h.AddSalary)
	salaryGroup.POST("/add/bulk", h.BulkAddSalaries)
	salaryGroup.GET("/all", h.GetAllSalary)
	salaryGroup.PUT("/update/:id", h.UpdateSalary)
	salaryGroup.DELETE("/delete/:id", h.DeleteSalary)
	salaryGroup.DELETE("/delete/bulk", h.BulkDeleteSalaries)
}

func (h *SalaryHandler) AddSalary(c echo.Context) error {
	var payload entity.CreateEmployeeSalary

	if err := c.Bind(&payload); err != nil {
		log.Println("Error binding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}

	res, err := h.SalaryService.AddSalaryService(c.Request().Context(), payload)
	fmt.Println(err)
	return c.JSON(http.StatusAccepted, res)
}

func (h *SalaryHandler) BulkAddSalaries(c echo.Context) error {

	var bulkRequest entity.BulkCreateEmployeeSalary

	if err := c.Bind(&bulkRequest); err != nil {
		log.Println("Error binding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}

	//fmt.Println("Binding:", bulkRequest)
	// Extract the salaries from the bulk request
	employee_salaries := bulkRequest.BulkSalaries

	// Call the service layer to handle bulk addition
	inserted_salaries, err := h.SalaryService.BulkAddSalaryService(c.Request().Context(), employee_salaries)
	if err != nil {
		log.Println("Error handling bulk addition:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return c.JSON(http.StatusAccepted, inserted_salaries)
}

func (h *SalaryHandler) GetAllSalary(c echo.Context) error {
	res, err := h.SalaryService.GetAllSalaryService(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusAccepted, res)
}

func (h *SalaryHandler) UpdateSalary(c echo.Context) error {
	var payload entity.CreateEmployeeSalary

	id := c.Param("id")

	if err := c.Bind(&payload); err != nil {
		log.Println("Error binding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}

	res, err := h.SalaryService.UpdateSalaryService(c.Request().Context(), payload, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}
	return c.JSON(http.StatusAccepted, res)
}

func (h *SalaryHandler) DeleteSalary(c echo.Context) error {
	var payload entity.EmployeeSalary

	id := c.Param("id")

	err := h.SalaryService.DeleteSalaryService(c.Request().Context(), id, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Salary deletion successful"})
}

func (h *SalaryHandler) BulkDeleteSalaries(c echo.Context) error {

	var employee_salary_ids entity.BulkDeleteSalaries

	if err := c.Bind(&employee_salary_ids); err != nil {
		log.Println("Error binding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}

	err := h.SalaryService.BulkDeleteSalaryService(c.Request().Context(), employee_salary_ids)
	if err != nil {
		log.Println("Error handling bulk delete:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete salaries"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Bulk deletion successful"})
}
