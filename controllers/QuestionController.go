package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pollp/controllers/utils"
	"pollp/models"
	"pollp/services"
)

type IQuestionController interface {
	GetQuestions(c *gin.Context)
	GetQuestion(c *gin.Context)
	CreateQuestion(c *gin.Context)
}

type QuestionController struct {
	questionService services.IQuestionService
}

func NewQuestionController(service services.IQuestionService) *QuestionController {
	return &QuestionController{
		questionService: service,
	}
}

func (qc *QuestionController) GetQuestions(c *gin.Context) {
	result := qc.questionService.GetQuestions()
	c.IndentedJSON(http.StatusOK, result)
}

func (qc *QuestionController) GetQuestion(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	result, err := qc.questionService.GetQuestionById(id)

	if err != nil {
		stop := utils.HandleControllerError(err, c)
		if stop {
			return
		}
	}

	c.IndentedJSON(http.StatusOK, result)
}

func (qc *QuestionController) CreateQuestion(c *gin.Context) {
	var question models.Question
	err := c.BindJSON(&question)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := qc.questionService.CreateQuestion(question)

	if err != nil {
		stop := utils.HandleControllerError(err, c)
		if stop {
			return
		}
	}

	c.IndentedJSON(http.StatusCreated, result)
}
