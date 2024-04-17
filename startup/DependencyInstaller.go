package startup

import (
	"go.uber.org/dig"
	"log"
	"pollp/controllers"
	"pollp/dal"
	"pollp/gin"
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
			Constructor: dal.NewMongoDatastore,
			Token:       "MongoDatastore",
		},
		{
			Constructor: dal.NewMongoDBQuestionRepository,
			Interface:   new(dal.IQuestionRepository),
			Token:       "MongoDBQuestionRepository",
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
			Constructor: controllers.NewVoteController,
			Interface:   new(controllers.IVoteController),
			Token:       "VoteController",
		},
		{
			Constructor: gin.NewPollpServer,
			Interface:   new(gin.IPollpServer),
			Token:       "PollpServer",
		},
	}

	container := dig.New()
	for _, dep := range dependencies {
		var err error
		if dep.Interface != nil {
			err = container.Provide(dep.Constructor, dig.As(dep.Interface))
		} else {
			err = container.Provide(dep.Constructor)
		}

		handleError(err)
	}
	return container
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
