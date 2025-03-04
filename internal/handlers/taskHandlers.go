package handlers

import (
	"PetProject/internal/taskService"
	"PetProject/internal/web/tasks"
	"context"
	"errors"
	"gorm.io/gorm"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("invalid request body")
	}
	taskToCreate := taskService.Task{
		Task:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.Service.DeleteTaskByID(uint(request.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tasks.DeleteTasksId404Response{}, nil
		}
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("invalid request body")
	}

	taskUpdate := taskService.Task{
		Task:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}
	updatedTask, err := h.Service.UpdateTaskByID(uint(request.Id), taskUpdate)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tasks.PatchTasksId404Response{}, nil
		}
		return nil, err
	}
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}
