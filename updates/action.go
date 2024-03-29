package updates

import (
	"fmt"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/buttons"
	"github.com/realtemirov/projects/tgbot/helper/const/action"
	"github.com/realtemirov/projects/tgbot/helper/const/word"
	"github.com/realtemirov/projects/tgbot/model"
	"github.com/spf13/cast"
)

func (h *Handler) SetTodoText(m *tg.Message) bool {

	var (
		msg    tg.MessageConfig
		result bool
	)

	todo, err := h.srvc.TodoService.GetNoSet(m.Chat.ID)
	if err != nil {
		fmt.Println(err.Error())
		msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")
		h.bot.Send(msg)
		return false
	}

	if m.Caption != "" {
		m.Text = m.Caption
	}
	if m.Text != "" {

		if err := h.srvc.TodoService.AddText(m.Chat.ID, m.Text); err != nil {
			fmt.Println(err.Error())
			msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")
			h.bot.Send(msg)
			return false
		}
	}

	//print photo
	// fmt.Println(m.Photo)

	if len(m.Photo) != 0 {
		if todo.Photo_URL != "" {
			msg = tg.NewMessage(m.Chat.ID, "have pictures")
			h.bot.Send(msg)
			return false
		}
		urls := m.Photo[len(m.Photo)-1].FileID
		if err := h.srvc.TodoService.AddPhotoURL(m.Chat.ID, urls); err != nil {
			fmt.Println(err.Error())
			msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")
			h.bot.Send(msg)
			return false
		}
	}
	todo, err = h.srvc.TodoService.GetNoSet(m.Chat.ID)
	if err != nil {
		fmt.Println(err.Error())
		msg = tg.NewMessage(m.Chat.ID, "Please retry, something is wrong")
		h.bot.Send(msg)
		return false
	}
	fmt.Println("ss", todo.Photo_URL, "ss")
	if len(strings.Split(todo.Photo_URL, "!")) > 1 {

		var medias []interface{}
		urls := strings.Split(todo.Photo_URL, "!")
		for i := 0; i < len(urls)-1; i++ {

			r := tg.NewInputMediaPhoto(tg.FileID(urls[i]))
			if i == 0 {
				r.Caption = todo.ToString()
				r.ParseMode = "HTML"
			}

			medias = append(medias, r)
		}

		phts := tg.MediaGroupConfig{
			ChatID:           m.Chat.ID,
			Media:            medias,
			ReplyToMessageID: m.MessageID,
		}
		fmt.Println(phts.Media...)
		h.bot.Send(phts)
		return true
	}

	msg = tg.NewMessage(m.Chat.ID, todo.ToString())
	msg.ReplyMarkup = buttons.Ok
	result = true

	msg.ReplyToMessageID = m.MessageID
	msg.ParseMode = "HTML"
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

func (h *Handler) CancelTodo(m *tg.Message) {
	if err := h.srvc.TodoService.DeleteSetIsFalse(m.Chat.ID); err != nil {
		fmt.Println(err.Error())
	}
	if err := h.srvc.UserService.SetAction(m.Chat.ID, action.EMPTY); err != nil {
		fmt.Println(err.Error())
	}

	msg := tg.NewMessage(m.Chat.ID, word.CANCEL)
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

func DeadlineTimes(h *Handler) ([]*model.Deadline, error) {
	notifs, err := h.srvc.TodoService.GetAllDeadlineTimes()
	if err != nil {
		fmt.Println(err.Error())
	}
	return notifs, err
}

func (h *Handler) SendTodo(id string) {
	todo, err := h.srvc.TodoService.GetByID(id)
	fmt.Println("send todo", todo.User_ID, todo.Text)
	if err != nil {
		fmt.Println(err.Error())
	}

	msg := tg.NewMessage(todo.User_ID, todo.ToString())
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = buttons.Menu
	h.bot.Send(msg)
}

func (h *Handler) SendMessage(id string, text, parseMode string) (interface{}, bool) {

	msg := tg.NewMessage(cast.ToInt64(id), text)
	fmt.Println("send message", id, text, parseMode)
	msg.ParseMode = parseMode
	msg.ReplyMarkup = buttons.Menu
	m, err := h.bot.Send(msg)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error(), false
	}

	return m, true
}

func (h *Handler) Ok(id int64) error {
	return h.srvc.UserService.SetAction(id, action.TODO_ADD)
}
