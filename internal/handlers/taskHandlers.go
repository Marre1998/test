package handlers

import (
	"PetProject/internal/taskService"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{service}
}

func (h *Handler) GetTaskHandler(c echo.Context) error {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *Handler) CreateTaskHandler(c echo.Context) error {
	var task taskService.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "invalid request body",
		})
	}
	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, createdTask)
}

func (h *Handler) UpdateTaskHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	var task taskService.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "invalid request body",
		})
	}
	updatedTask, err := h.Service.UpdateTaskByID(uint(id), task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, updatedTask)
}

func (h *Handler) DeleteTaskHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "Could not convert id to int",
		})
	}
	if err := h.Service.DeleteTaskByID(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "Task deleted",
	})
}
