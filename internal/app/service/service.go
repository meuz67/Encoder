package service

import (
	"net/http"
	"strconv"

	"main/internal/app/endpoint"

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
	var req RequestData
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shiftStr := c.Param("shift")
	shiftInt, err := strconv.Atoi(shiftStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var answer RequestData
	answer.Message = s.e.MoveText(shiftInt, req.Message)
	c.JSON(http.StatusOK, answer)
}
