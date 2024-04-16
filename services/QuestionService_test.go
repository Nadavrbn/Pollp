package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"pollp/models"
	"testing"
)

type mockIQuestionRepository struct {
	questions []models.Question
	err       error
}

func (m *mockIQuestionRepository) CreateQuestion(question models.Question) (models.Question, error) {
	if m.err != nil {
		return models.Question{}, m.err
	}
	m.questions = append(m.questions, question)
	return question, nil
}

func (m *mockIQuestionRepository) GetQuestions() []models.Question {
	return m.questions
}

func (m *mockIQuestionRepository) GetQuestion(id uint32) (models.Question, error) {
	for _, q := range m.questions {
		if q.PublicId == id {
			return q, nil
		}
	}
	return models.Question{}, errors.New("question not found")
}

func TestQuestionService_CreateQuestion(t *testing.T) {
	// Prepare
	repo := &mockIQuestionRepository{}
	service := NewQuestionService(repo)
	question := models.Question{PublicId: 1, Title: "What's your favorite color?"}

	// Execute
	createdQuestion, err := service.CreateQuestion(question)

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, question, createdQuestion)
	assert.Equal(t, []models.Question{question}, repo.questions)
}

func TestQuestionService_GetQuestions(t *testing.T) {
	// Prepare
	repo := &mockIQuestionRepository{}
	service := NewQuestionService(repo)
	questions := []models.Question{
		{PublicId: 1, Title: "Question 1"},
		{PublicId: 2, Title: "Question 2"},
	}
	repo.questions = questions

	// Execute
	result := service.GetQuestions()

	// Verify
	assert.Equal(t, questions, result)
}

func TestQuestionService_GetQuestionById(t *testing.T) {
	// Prepare
	repo := &mockIQuestionRepository{}
	service := NewQuestionService(repo)
	question := models.Question{PublicId: 1, Title: "What's your favorite color?"}
	repo.questions = append(repo.questions, question)

	// Execute
	result, err := service.GetQuestionById(1)

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, question, result)
}

func TestQuestionService_GetQuestionById_NotFound(t *testing.T) {
	// Prepare
	repo := &mockIQuestionRepository{}
	service := NewQuestionService(repo)

	// Execute
	_, err := service.GetQuestionById(1)

	// Verify
	assert.Error(t, err)
	assert.EqualError(t, err, "question not found")
}
