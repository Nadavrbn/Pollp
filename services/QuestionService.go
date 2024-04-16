package services

import (
	"github.com/oklog/ulid/v2"
	"pollp/dal"
	"pollp/models"
)

type QuestionService struct {
	IQuestionRepository dal.IQuestionRepository
}

func NewQuestionService(IQuestionRepository dal.IQuestionRepository) *QuestionService {
	return &QuestionService{IQuestionRepository: IQuestionRepository}
}

func (s *QuestionService) CreateQuestion(question models.Question) (models.Question, error) {
	question.PublicId = ulid.Make().String()
	for i := 0; i < len(question.Responses); i++ {
		question.Responses[i].PublicId = ulid.Make().String()
	}
	return s.IQuestionRepository.CreateQuestion(question)
}

func (s *QuestionService) GetQuestions() []models.Question {
	return s.IQuestionRepository.GetQuestions()
}

func (s *QuestionService) GetQuestionById(id string) (models.Question, error) {
	return s.IQuestionRepository.GetQuestion(id)
}
