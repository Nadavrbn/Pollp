package dal

import "pollp/models"

type IQuestionRepository interface {
	CreateQuestion(question models.Question) (models.Question, error)
	GetQuestions() []models.Question
	GetQuestion(id string) (models.Question, error)
}
