package services

import (
	"github.com/oklog/ulid/v2"
	"pollp/dal"
	"pollp/models"
)

type QuestionService struct {
	QuestionRepository dal.IQuestionRepository
}

func NewQuestionService(IQuestionRepository dal.IQuestionRepository) *QuestionService {
	return &QuestionService{QuestionRepository: IQuestionRepository}
}

func (s *QuestionService) CreateQuestion(question models.Question) (models.Question, error) {
	question.PublicId = ulid.Make().String()
	for i := 0; i < len(question.Responses); i++ {
		question.Responses[i].PublicId = ulid.Make().String()
	}
	return s.QuestionRepository.CreateQuestion(question)
}

func (s *QuestionService) GetQuestions() []models.Question {
	return s.QuestionRepository.GetQuestions()
}

func (s *QuestionService) GetQuestionById(id string) (models.Question, error) {
	return s.QuestionRepository.GetQuestion(id)
}
func (s *QuestionService) AddVote(questionId string, answerId string) (models.Question, error) {
	return s.QuestionRepository.AddVote(questionId, answerId)
}
func (s *QuestionService) RemoveVote(questionId string, answerId string) (models.Question, error) {
	return s.QuestionRepository.RemoveVote(questionId, answerId)
}
