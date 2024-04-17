package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pollp/controllers/utils"
	"pollp/services"
)

type IVoteController interface {
	AddVote(c *gin.Context)
	RemoveVote(c *gin.Context)
}

type VoteController struct {
	questionService services.IQuestionService
}

func NewVoteController(service services.IQuestionService) *VoteController {
	return &VoteController{
		questionService: service,
	}
}

func (vc *VoteController) AddVote(c *gin.Context) {
	questionId := c.Param("questionId")
	answerId := c.Param("answerId")
	if questionId == "" || answerId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	question, err := vc.questionService.AddVote(questionId, answerId)
	if err != nil {
		stop := utils.HandleControllerError(err, c)
		if stop {
			return
		}
	}
	c.IndentedJSON(http.StatusOK, question)
}

func (vc *VoteController) RemoveVote(c *gin.Context) {
	questionId := c.Param("questionId")
	answerId := c.Param("answerId")
	if questionId == "" || answerId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	question, err := vc.questionService.RemoveVote(questionId, answerId)
	if err != nil {
		stop := utils.HandleControllerError(err, c)
		if stop {
			return
		}
	}
	c.IndentedJSON(http.StatusOK, question)
}
