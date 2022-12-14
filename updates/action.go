package updates

import (
	"fmt"
	"strconv"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/buttons"
	"github.com/realtemirov/projects/tgbot/helper/const/action"
	"github.com/realtemirov/projects/tgbot/helper/convert"
	"github.com/realtemirov/projects/tgbot/model"
)

func SetTodoTitle(h *Handler, m *tg.Message) bool {

	var (
		todo   model.Todo
		msg    tg.MessageConfig
		result bool
	)

	key := fmt.Sprint(m.Chat.ID)

	todo.Title = m.Text

	res, err := convert.TodoToString(&todo)
	if err != nil {
		fmt.Println(err.Error())
		msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")

	} else {
		h.rds.Set("todo-"+key, res)
		h.rds.Set(key, action.TODO_ADD)
		msg = tg.NewMessage(m.Chat.ID, "Set your title")
		msg.ReplyMarkup = buttons.New_todo
		result = true
	}

	h.bot.Send(msg)
	return result
}

func SetTodoDescription(h *Handler, m *tg.Message) bool {

	var (
		msg tg.MessageConfig

		result bool
	)

	key := fmt.Sprint(m.Chat.ID)
	json, err := h.rds.Get("todo-" + key)
	if err != nil {
		fmt.Println(err.Error())
		msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")
		return false
	}

	todo, err := convert.StringToTodo(json)
	if err != nil {
		fmt.Println(err.Error())
		msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")
		return false
	}

	todo.Description = m.Text

	res, err := convert.TodoToString(todo)
	if err != nil {
		fmt.Println(err.Error())
		msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")

	} else {
		h.rds.Set("todo-"+key, res)
		h.rds.Set(key, action.TODO_ADD)
		msg = tg.NewMessage(m.Chat.ID, "Set your Description")
		msg.ReplyMarkup = buttons.New_todo
		result = true
	}

	h.bot.Send(msg)
	return result
}

func SetTodoNotification(h *Handler, m *tg.Message) bool {
	var (
		msg    tg.MessageConfig
		result bool
	)

	key := fmt.Sprint(m.Chat.ID)

	_, err := strconv.Atoi(m.Text)

	if err != nil {
		fmt.Println(err.Error())
		msg = tg.NewMessage(m.Chat.ID, "Please enter number")
		return false
	}

	err = h.rds.Set("notif-"+key, m.Text)

	if err != nil {
		fmt.Println(err.Error())
		msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")
		return false
	}
	result = true
	msg = tg.NewMessage(m.Chat.ID, "Set number,  enter date")
	msg.ReplyMarkup = buttons.Clock
	h.bot.Send(msg)
	return result
}


func Warning(h *Handler, m *tg.Message) {
	del := tg.NewDeleteMessage(m.Chat.ID, m.MessageID)
	h.bot.Request(del)

	msg := tg.NewMessage(m.Chat.ID, "Please save you new task")
	msg.ReplyMarkup = buttons.New_todo
	h.bot.Send(msg)
}

func SaveTodo(h *Handler, m *tg.Message) {

	key := fmt.Sprint(m.Chat.ID)
	res, err := h.rds.Get("todo-" + key)
	if err != nil {
		fmt.Println(err.Error())
	}

	todo, err := convert.StringToTodo(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	todo.User_ID = m.Chat.ID
	id, err := h.srvc.TodoService.Create(todo)
	if err != nil {
		fmt.Println(err.Error())
	}
	h.rds.Del("todo-" + key)
	h.rds.Set(key, action.EMPTY)

	msg := tg.NewMessage(m.Chat.ID, "Your task is saved, id: "+id)
	msg.ReplyMarkup = buttons.Todo
	h.bot.Send(msg)
}

func CancelNew(h *Handler, m *tg.Message) {
	key := fmt.Sprint(m.Chat.ID)
	h.rds.Del("todo-" + key)
	h.rds.Set(key, action.EMPTY)

	msg := tg.NewMessage(m.Chat.ID, "Your task is canceled")
	msg.ReplyMarkup = buttons.Todo
	h.bot.Send(msg)
}
