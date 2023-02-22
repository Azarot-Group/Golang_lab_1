package main

import (
	"lab1"
	"lab1/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(lab1.Server)

	if err := srv.Run("5000", handlers.InitRoutes()); err != nil {
		log.Fatal("server error:", err.Error())
	}
}
