package main

import (
	"log"
	"pollp/gin"
	"pollp/startup"
)

func main() {
	startup.SetupConfig()
	container := startup.BuildContainer()
	err := container.Invoke(gin.NewPollpServer)
	if err != nil {
		log.Fatal(err)
	}
}
