package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mongoapp.com/model"
)

//go:generate mockgen -destination=../mocks/repository/todo.go -package=repository mongoapp.com/repository TodoRepository
type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo model.Todo) (bool, error)
	GetAll() (model.Todos, error)
	Delete(id primitive.ObjectID) (bool, error)
}

func NewTodoRepositoryDB(collection *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{collection}
}

func (t *TodoRepositoryDB) Insert(todo model.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.InsertOne(ctx, todo)

	if result.InsertedID == nil || err != nil {
		errors.New("Failed to add Todo")
		return false, err
	}

	return true, nil
}

func (t *TodoRepositoryDB) GetAll() (model.Todos, error) {
	var todo model.Todo
	var todos []model.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatalln(err)
			return nil, err
		}

		todos = append(todos, todo)
	}
	return todos, nil

}

func (t *TodoRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.DeleteOne(ctx, id)

	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}

	return true, nil
}
