package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func GetHandler(c echo.Context) error {
	var messages []Message

	if err := DB.Find(&messages).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "Failed to get all messages",
		})
	}
	return c.JSON(http.StatusOK, &messages)
}

func PostHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not bind message",
		})
	}

	if err := DB.Create(&message).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "Failed to create message",
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "Message successfully sent",
	})
}

func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not convert id to int",
		})
	}
	var updatedMessage Message
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not bind message",
		})
	}

	if err := DB.Model(&Message{}).Where("id = ?", id).Updates(updatedMessage).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "Failed to update message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "Message successfully sent",
	})

}

func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not convert id to int",
		})
	}

	if err := DB.Delete(&Message{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "Failed to delete message",
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "Message successfully deleted",
	})
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})
	
	e := echo.New()
	e.GET("/", GetHandler)
	e.POST("/", PostHandler)
	e.PATCH("/:id", PatchHandler)
	e.DELETE("/:id", DeleteHandler)
	e.Start(":8080")
}
