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
	salaryGroup.GET("/all", h.GetAllSalary)
	salaryGroup.PUT("/update/:id", h.UpdateSalary)
	salaryGroup.DELETE("/delete/:id", h.DeleteSalary)
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
	return err
}
