package updates

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/buttons"
	"github.com/realtemirov/projects/tgbot/helper/const/action"
	"github.com/realtemirov/projects/tgbot/helper/convert"
	"github.com/realtemirov/projects/tgbot/model"
)

func (h *Handler) SetTodoTitle(m *tg.Message) bool {

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

func (h *Handler) SetTodoDescription(m *tg.Message) bool {

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
	fmt.Println("Done")
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

func (h *Handler) SetTodoNotification(m *tg.Message) bool {
	var (
		msg    tg.MessageConfig
		result bool
	)

	key := fmt.Sprint(m.Chat.ID)

	err := h.rds.Set("notif-"+key, m.Text)

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

func (h *Handler) SaveTodo(m *tg.Message, id string) {

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
	todo.Base.ID = id
	n_id, err := h.srvc.TodoService.Create(todo)
	if err != nil {
		fmt.Println(err.Error())
	}
	h.rds.Del("todo-" + key)
	h.rds.Set(key, action.EMPTY)

	msg := tg.NewMessage(m.Chat.ID, "Your task is saved, id: "+n_id)
	msg.ReplyMarkup = buttons.Todo
	h.bot.Send(msg)
}

func (h *Handler) CancelNew(m *tg.Message) {
	key := fmt.Sprint(m.Chat.ID)
	h.rds.Del("todo-" + key)
	h.rds.Set(key, action.EMPTY)

	msg := tg.NewMessage(m.Chat.ID, "Your task is canceled")
	msg.ReplyMarkup = buttons.Todo
	h.bot.Send(msg)
}

func NotificationTimes(h *Handler) ([]*model.Notification, error) {
	notifs, err := h.srvc.NotifService.GetAll()

	if err != nil {
		fmt.Println(err.Error())
	}

	return notifs, err
}

func SendTodo(h *Handler, id string) {
	fmt.Println("HEREEEEE")
	todo, err := h.srvc.TodoService.GetByID(id)
	fmt.Println("TODO", todo)
	if err != nil {
		fmt.Println(err.Error())
	}

	msg := tg.NewMessage(todo.User_ID, todo.ToString())
	msg.ParseMode = "markdown"
	msg.ReplyMarkup = buttons.Todo
	h.bot.Send(msg)
}
