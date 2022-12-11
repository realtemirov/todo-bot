package convert

import (
	"encoding/json"

	"github.com/realtemirov/projects/tgbot/model"
)

func StringToTodo(key string) (*model.Todo, error) {
	var todo model.Todo

	err := json.Unmarshal([]byte(key), &todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func TodoToString(todo *model.Todo) (string, error) {
	res, err := json.Marshal(todo)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
