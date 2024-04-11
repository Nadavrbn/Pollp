package repositories

import (
	"fmt"
	"github.com/samber/lo"
	"pollp/models"
)

type InMemoryIQuestionRepository struct {
	questions []models.Question
}

func NewInMemoryIQuestionRepository() *InMemoryIQuestionRepository {
	return &InMemoryIQuestionRepository{
		questions: []models.Question{
			{Id: 1, Title: "Where should we go for lunch?"},
			{Id: 2, Title: "What programming language should we use for the backend?"},
			{Id: 3, Title: "When can you go to laser tag?"},
		},
	}
}

func (r *InMemoryIQuestionRepository) CreateQuestion(question models.Question) (models.Question, error) {
	matches := lo.Filter(r.questions, func(x models.Question, index int) bool {
		return x.Id == question.Id
	})
	if len(matches) > 0 {
		return models.Question{}, fmt.Errorf("a question with id %d already exists", question.Id)
	}
	r.questions = append(r.questions, question)
	return question, nil
}

func (r *InMemoryIQuestionRepository) GetQuestions() []models.Question {
	return r.questions
}

func (r *InMemoryIQuestionRepository) GetQuestion(id uint32) (models.Question, error) {
	matches := lo.Filter(r.questions, func(x models.Question, index int) bool {
		return x.Id == id
	})
	if len(matches) > 1 {
		return models.Question{}, fmt.Errorf("there is more than 1 question with id %d", id)
	} else if len(matches) == 0 {
		return models.Question{}, fmt.Errorf("question with id %d not found", id)
	} else {
		return matches[0], nil
	}
}
