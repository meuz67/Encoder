package main

import (
	"main/internal/app/pkg"
	"main/internal/app/service"
)

func main() {
	s := service.NewService()
	a := pkg.New(s)
	a.Run()
}
