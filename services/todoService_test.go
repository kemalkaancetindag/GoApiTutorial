package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mongoapp.com/mocks/repository"
	"mongoapp.com/model"
)

var mockRepo *repository.MockTodoRepository
var service TodoService

var FakeData = []model.Todo{
	{
		Title:   "Title1",
		Content: "Content1",
	},
	{
		Title:   "Title2",
		Content: "Content2",
	},
	{
		Title:   "Title3",
		Content: "Content3",
	},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepo = repository.NewMockTodoRepository(ct)
	service = NewTodoService(mockRepo)

	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestDefaultTodoService_GetAll(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().GetAll().Return(FakeData, nil)
	result, err := service.GetAll()

	assert.NotEmpty(t, result)

	if err != nil {
		t.Error(err)
	}
}
