package pkg

import (
	"main/internal/app/service"

	"github.com/gin-gonic/gin"
)

type App struct {
	s service.Service
}

func New(s service.Service) *App {
	return &App{
		s: s,
	}
}
func (a *App) Run() {
	router := gin.Default()
	router.GET("/cipther/:shift", a.s.Shift)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
