package main

import (
	"PetProject/internal/database"
	"PetProject/internal/handlers"
	"PetProject/internal/taskService"
	"PetProject/internal/userService"
	"PetProject/internal/web/tasks"
	"PetProject/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()

	tasksRepo := taskService.NewTaskRepository(database.DB)
	usersRepo := userService.NewUserRepository(database.DB)
	tasksService := taskService.NewTaskService(tasksRepo)
	usersService := userService.NewUserService(usersRepo)

	tasksHandler := handlers.NewHandler(tasksService)
	usersHandler := handlers.NewUserHandler(usersService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)

	tasks.RegisterHandlers(e, strictHandler)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
