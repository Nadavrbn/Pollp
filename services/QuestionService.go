package services

import (
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
	return s.IQuestionRepository.CreateQuestion(question)
}

func (s *QuestionService) GetQuestions() []models.Question {
	return s.IQuestionRepository.GetQuestions()
}

func (s *QuestionService) GetQuestionById(id uint32) (models.Question, error) {
	return s.IQuestionRepository.GetQuestion(id)
}
