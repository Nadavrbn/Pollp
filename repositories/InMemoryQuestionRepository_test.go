package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pollp/models"
)

func TestInMemoryIQuestionRepository_CreateQuestion(t *testing.T) {
	repo := NewInMemoryIQuestionRepository()

	newQuestion := models.Question{Id: 5, Title: "How much wood would a woodchuck chuck if a woodchuck could chuck wood?"}
	createdQuestion, err := repo.CreateQuestion(newQuestion)

	assert.NoError(t, err, "Error should be nil")

	assert.Equal(t, newQuestion, createdQuestion, "Created question should match expected question")
}

func TestInMemoryIQuestionRepository_CreateQuestion_IdAlreadyExists(t *testing.T) {
	repo := NewInMemoryIQuestionRepository()

	newQuestion := models.Question{Id: 2, Title: "How much wood would a woodchuck chuck if a woodchuck could chuck wood?"}
	createdQuestion, err := repo.CreateQuestion(newQuestion)

	assert.Error(t, err, "An error should occur")
	assert.EqualError(t, err, "a question with id 2 already exists", "Error message should indicate id already exists")

	assert.Empty(t, createdQuestion, "Created question should match expected question")
}

func TestInMemoryIQuestionRepository_GetQuestions(t *testing.T) {
	repo := NewInMemoryIQuestionRepository()

	questions := repo.GetQuestions()

	assert.Len(t, questions, 3, "There should be 3 questions in the repository")
}

func TestInMemoryIQuestionRepository_GetQuestion(t *testing.T) {
	repo := NewInMemoryIQuestionRepository()

	expectedQuestion := models.Question{Id: 1, Title: "Where should we go for lunch?"}
	question, err := repo.GetQuestion(1)

	assert.NoError(t, err, "Error should be nil")

	assert.Equal(t, expectedQuestion, question, "Retrieved question should match expected question")
}

func TestInMemoryIQuestionRepository_GetQuestion_NotFound(t *testing.T) {
	repo := NewInMemoryIQuestionRepository()

	question, err := repo.GetQuestion(100)

	assert.Error(t, err, "An error should occur")
	assert.EqualError(t, err, "question with id 100 not found", "Error message should indicate question not found")

	assert.Empty(t, question)
}
