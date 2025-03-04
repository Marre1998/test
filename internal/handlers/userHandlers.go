package handlers

import (
	"PetProject/internal/userService"
	"PetProject/internal/web/users"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserHandler struct {
	Service *userService.UserService
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:    &usr.ID,
			Email: &usr.Email,
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("invalid request body")
	}
	userToCreate := userService.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}
	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:    &createdUser.ID,
		Email: &createdUser.Email,
	}
	return response, nil
}

func (h *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.Service.DeleteUserByID(uint(request.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.DeleteUsersId404Response{}, nil
		}
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}

func (h *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("invalid request body")
	}

	userUpdate := userService.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}
	updatedUser, err := h.Service.UpdateUserByID(uint(request.Id), userUpdate)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.PatchUsersId404Response{}, nil
		}
		return nil, err
	}
	response := users.PatchUsersId200JSONResponse{
		Id:    &updatedUser.ID,
		Email: &updatedUser.Email,
	}
	return response, nil
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}
