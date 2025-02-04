package main

import (
	"PetProject/internal/database"
	"PetProject/internal/handlers"
	"PetProject/internal/taskService"
	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewTaskService(repo)
	handler := handlers.NewHandler(service)

	e := echo.New()
	e.GET("/", handler.GetTaskHandler)
	e.POST("/", handler.CreateTaskHandler)
	e.PATCH("/:id", handler.UpdateTaskHandler)
	e.DELETE("/:id", handler.DeleteTaskHandler)
	e.Logger.Fatal(e.Start(":8080"))
	e.Start("8080")

}
