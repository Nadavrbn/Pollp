package services

import (
	"pollp/models"
	"pollp/repositories"
)

type QuestionService struct {
	IQuestionRepository repositories.IQuestionRepository
}

func NewQuestionService(IQuestionRepository repositories.IQuestionRepository) *QuestionService {
	return &QuestionService{IQuestionRepository: IQuestionRepository}
}

func (s *QuestionService) CreateQuestion(question models.Question) (models.Question, error) {
	return s.IQuestionRepository.CreateQuestion(question)
}

func (s *QuestionService) GetQuestions() []models.Question {
	return s.IQuestionRepository.GetQuestions()
}

func (s *QuestionService) GetQuestionById(id uint32) (models.Question, error) {
	return s.IQuestionRepository.GetQuestion(id)
}
