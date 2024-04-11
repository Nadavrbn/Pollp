package main

import (
	"log"
	"pollp/gin"
)

func main() {
	container := BuildContainer()
	err := container.Invoke(gin.NewPollpServer)
	if err != nil {
		log.Fatal(err)
	}
}
