package main

import (
	"go.uber.org/dig"
	"log"
	"pollp/controllers"
	"pollp/gin"
	"pollp/repositories"
	"pollp/services"
)

type Dependency struct {
	Constructor interface{}
	Interface   interface{}
	Token       string
}

func BuildContainer() *dig.Container {
	dependencies := []Dependency{
		{
			Constructor: repositories.NewInMemoryIQuestionRepository,
			Interface:   new(repositories.IQuestionRepository),
			Token:       "InMemoryQuestionRepository",
		},
		{
			Constructor: services.NewQuestionService,
			Interface:   new(services.IQuestionService),
			Token:       "QuestionService",
		},
		{
			Constructor: controllers.NewQuestionController,
			Interface:   new(controllers.IQuestionController),
			Token:       "QuestionController",
		},
		{
			Constructor: gin.NewPollpServer,
			Interface:   new(gin.IPollpServer),
			Token:       "PollpServer",
		},
	}

	container := dig.New()
	for _, dep := range dependencies {
		err := container.Provide(dep.Constructor, dig.As(dep.Interface))
		handleError(err)
	}
	return container
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
