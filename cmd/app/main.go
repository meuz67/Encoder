package main

import (
	"log"

	"main/internal/app/pkg"
	"main/internal/app/service"
)

func main() {
	s, err := service.NewService()
	if err != nil {
		log.Fatal(err)
	}

	a := pkg.New(s)
	a.Run()
}
