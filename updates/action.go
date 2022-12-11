package updates

import (
	"fmt"

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

func Warning(h *Handler, m *tg.Message) {
	del := tg.NewDeleteMessage(m.Chat.ID, m.MessageID)
	h.bot.Request(del)

	msg := tg.NewMessage(m.Chat.ID, "Please save you new task")
	msg.ReplyMarkup = buttons.New_todo
	h.bot.Send(msg)
}

func Notification(h *Handler, m *tg.Message) {

	msg := tg.NewMessage(m.Chat.ID, "Telegram orqali eslatma olasizmi?")
	msg.ReplyMarkup = buttons.Notification
	h.bot.Send(msg)
}

func SaveTodo(h *Handler, m *tg.Message) {

	msg := tg.NewMessage(m.Chat.ID, "Telegram orqali eslatma olasizmi?")
	msg.ReplyMarkup = buttons.Notification
	h.bot.Send(msg)

	/*key := fmt.Sprint(m.Chat.ID)
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
	h.bot.Send(msg)*/
}
