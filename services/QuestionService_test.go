package services_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mocks "pollp/mocks"
	"pollp/models"
	"pollp/services"
)

func TestQuestionService_CreateQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIQuestionRepository(ctrl)
	service := services.NewQuestionService(mockRepo)

	question := models.Question{Title: "What's your favorite color?"}
	mockRepo.EXPECT().CreateQuestion(gomock.Any()).DoAndReturn(func(q models.Question) (models.Question, error) {
		return q, nil
	})
	createdQuestion, err := service.CreateQuestion(question)

	assert.NoError(t, err)
	assert.Regexp(t, "[0-7][0-9A-HJKMNP-TV-Z]{25}", createdQuestion.PublicId, "Question should have generated public id in ULID format")
	question.PublicId = createdQuestion.PublicId
	assert.Equal(t, question, createdQuestion)
}

func TestQuestionService_GetQuestions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIQuestionRepository(ctrl)
	service := services.NewQuestionService(mockRepo)

	questions := []models.Question{
		{PublicId: "1AB2", Title: "Question 1"},
		{PublicId: "3CD4", Title: "Question 2"},
	}
	mockRepo.EXPECT().GetQuestions().Return(questions)

	result := service.GetQuestions()

	assert.Equal(t, questions, result)
}

func TestQuestionService_GetQuestionById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIQuestionRepository(ctrl)
	service := services.NewQuestionService(mockRepo)

	question := models.Question{PublicId: "1AB2", Title: "What's your favorite color?"}
	mockRepo.EXPECT().GetQuestion("1AB2").Return(question, nil)

	result, err := service.GetQuestionById("1AB2")

	assert.NoError(t, err)
	assert.Equal(t, question, result)
}

func TestQuestionService_GetQuestionById_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIQuestionRepository(ctrl)
	service := services.NewQuestionService(mockRepo)

	mockRepo.EXPECT().GetQuestion("1AZ2").Return(models.Question{}, errors.New("question not found"))

	_, err := service.GetQuestionById("1AZ2")

	assert.Error(t, err)
	assert.EqualError(t, err, "question not found")
}

func TestQuestionService_AddVote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIQuestionRepository(ctrl)
	service := services.NewQuestionService(mockRepo)

	questionID := "01HVKXMHJ08YX357CJDC11A41B"
	answerID := "01HVKXMHJ08YX357CJDC8WVYE1"
	expectedQuestion := models.Question{ /* construct expected question */ }
	mockRepo.EXPECT().AddVote(questionID, answerID).DoAndReturn(func(qID, aID string) (models.Question, error) {
		return expectedQuestion, nil
	})

	result, err := service.AddVote(questionID, answerID)

	assert.NoError(t, err)
	assert.Equal(t, expectedQuestion, result)
}

func TestQuestionService_RemoveVote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIQuestionRepository(ctrl)
	service := services.NewQuestionService(mockRepo)

	questionID := "01HVKXMHJ08YX357CJDC11A41B"
	answerID := "01HVKXMHJ08YX357CJDC8WVYE1"
	expectedQuestion := models.Question{ /* construct expected question */ }
	mockRepo.EXPECT().RemoveVote(questionID, answerID).DoAndReturn(func(qID, aID string) (models.Question, error) {
		return expectedQuestion, nil
	})

	result, err := service.RemoveVote(questionID, answerID)

	assert.NoError(t, err)
	assert.Equal(t, expectedQuestion, result)
}
