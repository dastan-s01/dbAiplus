package handlers

import (
	"dbAiplus/internal/app/di"
	"dbAiplus/internal/app/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {
	di *di.DI
}

func NewHandler(di *di.DI) *Handler {
	return &Handler{di: di}
}
func (h *Handler) InitRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/employees", h.CreateEmployeeHandler)
	}
}
func (h *Handler) CreateEmployeeHandler(c *gin.Context) {
	var employee models.Employee

	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	err := h.di.EmployeeUseCase.AddEmployee(c.Request.Context(), employee)
	if err != nil {
		log.Println("Ошибка при добавлении сотрудника:", err) // Логируем ошибку
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении сотрудника"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Сотрудник успешно добавлен"})
}
