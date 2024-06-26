package dal

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
			{PublicId: "1", Title: "Where should we go for lunch?"},
			{PublicId: "2", Title: "What programming language should we use for the backend?"},
			{PublicId: "3", Title: "When can you go to laser tag?"},
		},
	}
}

func (r *InMemoryIQuestionRepository) CreateQuestion(question models.Question) (models.Question, error) {
	matches := lo.Filter(r.questions, func(x models.Question, index int) bool {
		return x.PublicId == question.PublicId
	})
	if len(matches) > 0 {
		return models.Question{}, fmt.Errorf("a question with id %s already exists", question.PublicId)
	}
	r.questions = append(r.questions, question)
	return question, nil
}

func (r *InMemoryIQuestionRepository) GetQuestions() []models.Question {
	return r.questions
}

func (r *InMemoryIQuestionRepository) GetQuestion(id string) (models.Question, error) {
	matches := lo.Filter(r.questions, func(x models.Question, index int) bool {
		return x.PublicId == id
	})
	if len(matches) > 1 {
		return models.Question{}, fmt.Errorf("there is more than 1 question with id %s", id)
	} else if len(matches) == 0 {
		return models.Question{}, fmt.Errorf("question with id %s not found", id)
	} else {
		return matches[0], nil
	}
}
