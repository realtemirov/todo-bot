package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/realtemirov/projects/tgbot/model"
	"github.com/realtemirov/projects/tgbot/storage"
)

type todoService struct {
	todoRepo storage.TodoI
}

func NewTodoRepository(todoRepo storage.TodoI) *todoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}

func (t *todoService) Create(todo *model.Todo) (string, error) {
	return t.todoRepo.Create(&model.Todo{
		Base: model.Base{
			ID:        uuid.New().String(),
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		},
		User_ID:     todo.User_ID,
		Title:       todo.Title,
		Description: todo.Description,
		Photo_URL:   todo.Photo_URL,
		File_URL:    todo.File_URL,
		Deadline:    todo.Deadline,
		IsDone:      false,
	})
}

func (t *todoService) GetByID(id string) (*model.Todo, error) {
	return t.todoRepo.GetByID(id)
}

func (t *todoService) GetAll(id int64) ([]*model.Todo, error) {
	return t.todoRepo.GetAllByUserID(id)
}

func (t *todoService) Update(todo *model.Todo) (*model.Todo, error) {
	time := time.Now()
	return t.todoRepo.Update(&model.Todo{

		Base: model.Base{
			ID:        todo.ID,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: &time,
			DeletedAt: nil,
		},
		User_ID:     todo.User_ID,
		Title:       todo.Title,
		Description: todo.Description,
		Photo_URL:   todo.Photo_URL,
		File_URL:    todo.File_URL,
		Deadline:    todo.Deadline,
		IsDone:      false,
	})
}

func (t *todoService) Delete(id string) error {
	return t.todoRepo.Delete(id)
}
