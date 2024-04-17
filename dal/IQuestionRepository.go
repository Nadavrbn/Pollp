package dal

import "pollp/models"

type IQuestionRepository interface {
	CreateQuestion(question models.Question) (models.Question, error)
	GetQuestions() []models.Question
	GetQuestion(id string) (models.Question, error)
	AddVote(questionId string, answerId string) (models.Question, error)
	RemoveVote(questionId string, answerId string) (models.Question, error)
}
