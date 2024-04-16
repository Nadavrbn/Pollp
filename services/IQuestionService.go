package services

import "pollp/models"

type IQuestionService interface {
	CreateQuestion(question models.Question) (models.Question, error)
	GetQuestions() []models.Question
	GetQuestionById(id uint32) (models.Question, error)
	GetQuestionById(id string) (models.Question, error)
}
