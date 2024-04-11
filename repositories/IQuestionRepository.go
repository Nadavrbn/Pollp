package repositories

import "pollp/models"

type IQuestionRepository interface {
	CreateQuestion(question models.Question) (models.Question, error)
	GetQuestions() []models.Question
	GetQuestion(id uint32) (models.Question, error)
}
