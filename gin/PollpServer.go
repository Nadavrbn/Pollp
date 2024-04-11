package gin

import (
	"github.com/gin-gonic/gin"
	"pollp/controllers"
)

type PollpServer struct {
	QuestionController controllers.IQuestionController
}

func NewPollpServer(questionController controllers.IQuestionController) *PollpServer {
	server := &PollpServer{
		QuestionController: questionController,
	}

	engine := gin.Default()
	engine.GET("/questions", server.QuestionController.GetQuestions)
	engine.GET("/questions/:id", server.QuestionController.GetQuestion)
	engine.POST("/questions", server.QuestionController.CreateQuestion)

	engine.Run()
	return server
}
