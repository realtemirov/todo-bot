package updates

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/buttons"
	"github.com/realtemirov/projects/tgbot/helper/const/action"
	"github.com/realtemirov/projects/tgbot/model"
)

func (h *Handler) SetTodoTitle(m *tg.Message) bool {

	var (
		todo   *model.Todo
		msg    tg.MessageConfig
		result bool
	)

	todo = &model.Todo{
		User_ID: m.Chat.ID,
		Title:   m.Text,
		Is_Set:  false,
	}

	if _, err := h.srvc.TodoService.Create(todo); err != nil {
		fmt.Println(err.Error())
		msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")

	}
	if err := h.srvc.UserService.SetAction(m.Chat.ID, action.TODO_ADD); err != nil {
		fmt.Println(err.Error())
		msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")
	} else {
		msg = tg.NewMessage(m.Chat.ID, "Set your title")
		msg.ReplyMarkup = buttons.New_todo
		result = true
	}

	msg.ReplyToMessageID = m.MessageID
	h.bot.Send(msg)
	return result
}

func (h *Handler) SetTodoDescription(m *tg.Message) bool {

	var (
		msg    tg.MessageConfig
		result bool
	)

	if err := h.srvc.TodoService.AddDescription(m.Chat.ID, m.Text); err != nil {
		fmt.Println(err.Error())
		msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")
	}
	if err := h.srvc.UserService.SetAction(m.Chat.ID, action.TODO_ADD); err != nil {
		fmt.Println(err.Error())
		msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")
	} else {
		msg = tg.NewMessage(m.Chat.ID, "Set your Description")
		msg.ReplyMarkup = buttons.New_todo
		result = true
	}

	msg.ReplyToMessageID = m.MessageID
	h.bot.Send(msg)
	return result
}

// func (h *Handler) SetTodoNotification(m *tg.Message) bool {
// 	var (
// 		msg    tg.MessageConfig
// 		result bool
// 	)

// 	key := fmt.Sprint(m.Chat.ID)

// 	err := h.rds.Set("notif-"+key, m.Text)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")
// 		return false
// 	}
// 	result = true
// 	msg = tg.NewMessage(m.Chat.ID, "Set number,  enter date")
// 	msg.ReplyMarkup = buttons.Clock
// 	h.bot.Send(msg)
// 	return result
// }

// func Warning(h *Handler, m *tg.Message) {
// 	del := tg.NewDeleteMessage(m.Chat.ID, m.MessageID)
// 	h.bot.Request(del)

// 	msg := tg.NewMessage(m.Chat.ID, "Please save you new task")
// 	msg.ReplyMarkup = buttons.New_todo
// 	h.bot.Send(msg)
// }

// func (h *Handler) SaveTodo(m *tg.Message, id string) {

// 	key := fmt.Sprint(m.Chat.ID)
// 	res, err := h.rds.Get("todo-" + key)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	todo, err := convert.StringToTodo(res)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	todo.User_ID = m.Chat.ID
// 	todo.Base.ID = id
// 	n_id, err := h.srvc.TodoService.Create(todo)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	h.rds.Del("todo-" + key)
// 	h.rds.Set(key, action.EMPTY)

// 	msg := tg.NewMessage(m.Chat.ID, "Your task is saved, id: "+n_id)
// 	msg.ReplyMarkup = buttons.Todo
// 	h.bot.Send(msg)
// }

func (h *Handler) CancelNew(m *tg.Message) {
	if err := h.srvc.TodoService.DeleteSetIsFalse(m.Chat.ID); err != nil {
		fmt.Println(err.Error())
	}
	if err := h.srvc.UserService.SetAction(m.Chat.ID, action.EMPTY); err != nil {
		fmt.Println(err.Error())
	}

	msg := tg.NewMessage(m.Chat.ID, "Your task is canceled")
	msg.ReplyToMessageID = m.MessageID
	msg.ReplyMarkup = buttons.Menu
	h.bot.Send(msg)
}

func NotificationTimes(h *Handler) ([]*model.Notification, error) {
	notifs, err := h.srvc.TodoService.GetAllNotificationTimes()
	if err != nil {
		fmt.Println(err.Error())
	}
	return notifs, err
}

func SendTodo(h *Handler, id string) {
	todo, err := h.srvc.TodoService.GetByID(id)
	fmt.Println("send todo", todo.User_ID, todo.Title)
	if err != nil {
		fmt.Println(err.Error())
	}

	msg := tg.NewMessage(todo.User_ID, todo.ToString())
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = buttons.Todo
	h.bot.Send(msg)
}
