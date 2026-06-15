package service

import (
	"net/http"
	"strconv"

	"main/internal/app/endpoint"
	"main/internal/database"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	Message string `json:"message"`
}

type Service struct {
	e  *endpoint.Endpoint
	db *database.Database
}

func NewService() (*Service, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	return &Service{
		e:  endpoint.NewEndpoint(),
		db: db,
	}, nil
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
	if err := s.db.SaveMessage(c.Request.Context(), shiftInt, req.Message, answer.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, answer)
}
