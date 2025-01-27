package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var messages = make(map[int]Message)
var nextID = 1

func GetHandler(c echo.Context) error {
	var msgSlice []Message
	for _, msg := range messages {
		msgSlice = append(msgSlice, msg)
	}
	return c.JSON(http.StatusOK, &msgSlice)
}

func PostHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not bind message",
		})
	}
	message.ID = nextID
	nextID++

	messages[message.ID] = message
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

	if _, ok := messages[id]; !ok {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not find message",
		})
	}
	updatedMessage.ID = id
	messages[id] = updatedMessage

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
	if _, ok := messages[id]; !ok {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not find message",
		})
	}
	delete(messages, id)
	return c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "Message successfully deleted",
	})
}

func main() {
	e := echo.New()
	e.GET("/", GetHandler)
	e.POST("/", PostHandler)
	e.PATCH("/:id", PatchHandler)
	e.DELETE("/:id", DeleteHandler)
	e.Start(":8080")
}
