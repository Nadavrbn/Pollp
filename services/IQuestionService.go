package services

import "pollp/models"

type IQuestionService interface {
	CreateQuestion(question models.Question) (models.Question, error)
	GetQuestions() []models.Question
	GetQuestionById(id string) (models.Question, error)
	AddVote(questionId string, answerId string) (models.Question, error)
	RemoveVote(questionId string, answerId string) (models.Question, error)
}
