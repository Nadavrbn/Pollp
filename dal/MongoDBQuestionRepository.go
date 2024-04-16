package dal

import (
	"context"
	"errors"
	"pollp/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBQuestionRepository struct {
	mongo      *MongoDatastore
	collection *mongo.Collection
}

func NewMongoDBQuestionRepository(mongo *MongoDatastore) *MongoDBQuestionRepository {
	collection := mongo.DB.Collection("questions")
	return &MongoDBQuestionRepository{
		mongo:      mongo,
		collection: collection,
	}
}

// CreateQuestion inserts a new question into the MongoDB collection.
func (r *MongoDBQuestionRepository) CreateQuestion(question models.Question) (models.Question, error) {
	result, err := r.collection.InsertOne(context.TODO(), question)
	if err != nil {
		return models.Question{}, err
	}
	question.Id = result.InsertedID
	return question, nil
}

// GetQuestions retrieves all questions from the MongoDB collection.
func (r *MongoDBQuestionRepository) GetQuestions() []models.Question {
	var questions []models.Question

	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil //, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var question models.Question
		err := cursor.Decode(&question)
		if err != nil {
			return nil //, err
		}
		questions = append(questions, question)
	}

	//if err := cursor.Err(); err != nil {
	//	return nil, err
	//}

	return questions
}

// GetQuestion retrieves a question by its ID from the MongoDB collection.
func (r *MongoDBQuestionRepository) GetQuestion(id uint32) (models.Question, error) {
	var question models.Question

	filter := bson.M{"publicId": id}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&question)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Question{}, errors.New("question not found")
		}
		return models.Question{}, err
	}

	return question, nil
}
