package updates

import (
	"fmt"

	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/buttons"
	"github.com/realtemirov/projects/tgbot/helper"
	"github.com/realtemirov/projects/tgbot/helper/const/action"
	"github.com/realtemirov/projects/tgbot/helper/const/word"
	"github.com/realtemirov/projects/tgbot/model"
	"github.com/spf13/cast"
)

func (h *Handler) Profile(m *tg.Message) {

	user, err := h.srvc.UserService.GetByID(m.Chat.ID)
	if err != nil {
		fmt.Println(err.Error() + " in getme")
	}
	cnf := tg.UserProfilePhotosConfig{
		UserID: m.Chat.ID,
		Offset: 0,
		Limit:  1,
	}

	photos, err := h.bot.GetUserProfilePhotos(cnf)
	if err == nil && photos.TotalCount != 0 {
		p := tg.NewPhoto(user.ID, tg.FileID(photos.Photos[0][2].FileID))
		p.Caption = helper.UserToString(m.From, user)
		p.ParseMode = "HTML"
		h.bot.Send(p)
	} else {
		msg := tg.NewMessage(m.Chat.ID, helper.UserToString(m.From, user))
		msg.ParseMode = "HTML"
		h.bot.Send(msg)
	}
}

func (h *Handler) SingUp(m *tg.Message) {
	user := &model.User{
		ID:     m.Chat.ID,
		Action: action.EMPTY,
	}

	id, err := h.srvc.UserService.Add(user)
	text := ""
	if err != nil && id == 0 {
		if strings.Contains(err.Error(), "duplicate") {
			text = fmt.Sprintf("%s"+word.START, m.Chat.FirstName)
		} else {
			text = fmt.Sprintf("Error %s", err.Error())
		}

	} else {
		text = fmt.Sprintf("Salom, <b>%s</b>, sizning ID : <code>%v</code> \nKeyingi amalni tanlang: ", m.Chat.FirstName, id)
	}

	msg := tg.NewMessage(m.Chat.ID, text)
	msg.ReplyToMessageID = m.MessageID
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = buttons.Menu
	h.bot.Send(msg)
}

func (h *Handler) AddTodo(m *tg.Message) {

	text := word.TODO
	msg := tg.NewMessage(m.Chat.ID, text)

	act, err := h.srvc.UserService.GetAction(m.Chat.ID)
	if err != nil {
		fmt.Println("Err : " + err.Error())
	}
	if act == action.EMPTY {
		if err := h.srvc.UserService.SetAction(m.Chat.ID, action.TODO_NEW_TITLE); err != nil {
			fmt.Println("Err : " + err.Error())
		}
		_, err := h.srvc.TodoService.Create(&model.Todo{
			User_ID: m.Chat.ID,
		})
		if err != nil {
			fmt.Println("Err : " + err.Error())
		}
	}
	msg.ReplyToMessageID = m.MessageID
	msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
	h.bot.Send(msg)
}

// func (h *Handler) AddNotification(m *tg.Message) {

//		h.srvc.UserService.SetAction(m.Chat.ID, action.TODO_NEW_NOTIFICATION)
//		msg := tg.NewMessage(m.Chat.ID, "Vaqtni tanlang : ")
//		msg.ReplyToMessageID = m.MessageID
//		msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
//		msg.ReplyMarkup = buttons.Hour(ssssssssssssssssssssss)
//		h.bot.Send(msg)
//	}
func (h *Handler) DeadlineOrNotification(m *tg.Message) {
	h.srvc.UserService.SetAction(m.Chat.ID, action.TODO_NEW_TIME)
	msg := tg.NewMessage(m.Chat.ID, "Deadline Or Notification : ")
	msg.ReplyToMessageID = m.MessageID
	msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
	msg.ReplyMarkup = buttons.DeadlineOrNotification
	h.bot.Send(msg)
}
func (h *Handler) GetAllTodosByUserID(m *tg.Message, done bool, action string) bool {
	id := m.Chat.ID
	todos, err := h.srvc.TodoService.GetAllByUserID(id, done)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if len(todos) == 0 {
		return false
	}

	keyboard := helper.SliceToInlineKeyboard(todos, action)

	text := ""
	for i, v := range todos {
		text += cast.ToString(i+1) + ". " + v.ToString()
	}
	msg := tg.NewMessage(id, text)
	msg.ReplyMarkup = keyboard

	msg.ParseMode = "HTML"
	_, err = h.bot.Send(msg)
	if err != nil {
		fmt.Println(err.Error())
	}
	return true
}

func (h *Handler) TodoView(id, action string, m *tg.Message) bool {

	todo, err := h.srvc.TodoService.GetByID(id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if todo.Photo_URL != "" {

		var medias []interface{}
		urls := strings.Split(todo.Photo_URL, "!")
		r := tg.NewInputMediaPhoto(tg.FileID(urls[0]))
		r.Caption = todo.ToString()
		r.ParseMode = "HTML"
		medias = append(medias, r)

		for i := 0; i < len(urls)-1; i++ {
			if i == 0 {
				continue
			} else {
				r = tg.NewInputMediaPhoto(tg.FileID(urls[i]))
			}

			medias = append(medias, r)
		}

		phts := tg.MediaGroupConfig{
			ChatID: m.Chat.ID,
			Media:  medias,
		}

		msg, err := h.bot.SendMediaGroup(phts)
		if err != nil {
			fmt.Println("Err", err.Error())
		}

		edit := tg.NewEditMessageReplyMarkup(m.Chat.ID, msg[0].MessageID, buttons.Todo_view(id, action))
		//edit := tg.NewEditMessageCaption(msg[0].Chat.ID, msg[0].MessageID, "Hello") //
		_, err = h.bot.Send(edit)
		if err != nil {
			fmt.Println("Err55", err.Error())
			return false
		}

		return true
	}
	msg := tg.NewMessage(m.Chat.ID, todo.ToString())
	msg.ReplyMarkup = buttons.Todo_view(id, action)
	msg.ParseMode = "HTML"
	h.bot.Send(msg)
	return true
}

func (h *Handler) TodoDone(id, action string, m *tg.Message) bool {

	todo, err := h.srvc.TodoService.GetByID(id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if todo.Photo_URL != "" {

		var medias []interface{}
		urls := strings.Split(todo.Photo_URL, "!")
		r := tg.NewInputMediaPhoto(tg.FileID(urls[0]))
		r.Caption = todo.ToString()
		r.ParseMode = "HTML"
		medias = append(medias, r)

		for i := 0; i < len(urls)-1; i++ {
			if i == 0 {
				continue
			} else {
				r = tg.NewInputMediaPhoto(tg.FileID(urls[i]))
			}

			medias = append(medias, r)
		}

		phts := tg.MediaGroupConfig{
			ChatID: m.Chat.ID,
			Media:  medias,
		}

		msg, err := h.bot.SendMediaGroup(phts)
		if err != nil {
			fmt.Println("Err", err.Error())
		}

		edit := tg.NewEditMessageReplyMarkup(m.Chat.ID, msg[0].MessageID, buttons.Todo_Done(id, action))
		//edit := tg.NewEditMessageCaption(msg[0].Chat.ID, msg[0].MessageID, "Hello") //
		_, err = h.bot.Send(edit)
		if err != nil {
			fmt.Println("Err55", err.Error())
			return false
		}

		return true
	}
	msg := tg.NewMessage(m.Chat.ID, todo.ToString())
	msg.ReplyMarkup = buttons.Todo_Done(id, action)
	msg.ParseMode = "HTML"
	h.bot.Send(msg)
	return true
}
