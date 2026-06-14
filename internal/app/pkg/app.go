package pkg

import (
	"main/internal/app/service"

	"github.com/gin-gonic/gin"
)

type App struct {
	s *service.Service
}

func New(s *service.Service) *App {
	return &App{
		s: s,
	}
}

func (a *App) Run() {
	router := gin.Default()
	router.POST("/cipher/:shift", a.s.Shift)
	router.Run(":8080")
}
