package service

import (
	"main/internal/app/endpoint"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	Message string `json:"message"`
}

type Service struct {
	e *endpoint.Endpoint
}

func NewService() *Service {
	return &Service{
		e: endpoint.NewEndpoint(),
	}
}
func (s *Service) Shift(c *gin.Context) {
	var json RequestData
	if err := c.ShouldBindJSON(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	shiftStr := c.Param("shift")
	shiftInt, err := strconv.Atoi(shiftStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var answer RequestData
	answer.Message = s.e.MoveText(shiftInt, json.Message)
	c.JSON(http.StatusOK, answer)
}
