package main

import (
	"lab1"
	"lab1/handler"
	"log"
)

const port string = "5000"

func main() {
	handlers := new(handler.Handler)
	srv := new(lab1.Server)

	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatal("server error:", err.Error())
	}
}
