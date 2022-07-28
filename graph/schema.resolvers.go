package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	database "github.com/primozh/gin-graphql-postgres/db"
	"github.com/primozh/gin-graphql-postgres/graph/generated"
	"github.com/primozh/gin-graphql-postgres/graph/model"
)

// CreateQuestion is the resolver for the createQuestion field.
func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.QuestionInput) (*model.Question, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	fmt.Println("input", input.QuestionText, input.PubDate)
	question := model.Question{}
	question.QuestionText = input.QuestionText
	question.PubDate = input.PubDate
	db.Create(&question)
	return &question, nil
}

// CreateChoice is the resolver for the createChoice field.
func (r *mutationResolver) CreateChoice(ctx context.Context, input model.ChoiceInput) (*model.Choice, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	choice := model.Choice{}
	question := model.Question{}
	choice.QuestionID = input.QuestionID
	choice.ChoiceText = input.ChoiceText
	db.First(&question, choice.QuestionID)
	choice.Question = &question
	db.Create(&choice)
	return &choice, nil
}

// Questions is the resolver for the questions field.
func (r *queryResolver) Questions(ctx context.Context) ([]*model.Question, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	db.Find(&r.questions)
	for _, question := range r.questions {
		var choices []*model.Choice
		db.Where(&model.Choice{QuestionID: question.ID}).Find(&choices)
		question.Choices = choices
	}
	return r.questions, nil
}

// Choices is the resolver for the choices field.
func (r *queryResolver) Choices(ctx context.Context) ([]*model.Choice, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	db.Find(&r.choices)
	return r.choices, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
