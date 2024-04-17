package gin

import (
	"github.com/gin-gonic/gin"
	"pollp/controllers"
)

type PollpServer struct {
	QuestionController controllers.IQuestionController
	VoteController     controllers.IVoteController
}

func NewPollpServer(questionController controllers.IQuestionController, voteController controllers.IVoteController) *PollpServer {
	server := &PollpServer{
		QuestionController: questionController,
		VoteController:     voteController,
	}

	engine := gin.Default()
	engine.GET("/questions", server.QuestionController.GetQuestions)
	engine.GET("/questions/:id", server.QuestionController.GetQuestion)
	engine.POST("/questions", server.QuestionController.CreateQuestion)
	engine.POST("/questions/:questionId/votes/:answerId", server.VoteController.AddVote)
	engine.DELETE("/questions/:questionId/votes/:answerId", server.VoteController.RemoveVote)

	engine.Run()
	return server
}
