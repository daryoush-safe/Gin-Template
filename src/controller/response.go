package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type singleMessageResponse struct {
	Status  int         `json:"statusCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type multipleMessageResponse struct {
	Status   int                          `json:"statusCode"`
	Messages map[string]map[string]string `json:"messages"`
	Data     interface{}                  `json:"data"`
}

func Response[T string | map[string]map[string]string](c *gin.Context, statusCode int, message T, data interface{}) {
	statusText := http.StatusText(statusCode)
	if statusText == "" {
		panic("invalid status code!")
	}

	switch msg := any(message).(type) {
	case string:
		if msg == "" {
			msg = statusText
		}
		c.JSON(statusCode, singleMessageResponse{
			Status:  statusCode,
			Message: msg,
			Data:    data,
		})
	case map[string]map[string]string:
		c.JSON(statusCode, multipleMessageResponse{
			Status:   statusCode,
			Messages: msg,
			Data:     data,
		})
	}
}
