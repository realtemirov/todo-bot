package updates

import (
	"fmt"

	"github.com/realtemirov/projects/tgbot/model"
)

func (h *Handler) GetAllUsers() ([]*model.User, error) {
	users, err := h.srvc.UserService.GetAll()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return users, nil
}

func (h *Handler) GetAllTodos(id int64) ([]*model.Todo, error) {
	fmt.Println(id)
	todos, err := h.srvc.TodoService.GetAllByUserID(id, false)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return todos, nil
}
