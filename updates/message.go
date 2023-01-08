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
		p := tg.NewPhoto(user.ID, tg.FileID(photos.Photos[0][0].FileID))
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

func (h *Handler) Todo(m *tg.Message) {
	msg := tg.NewMessage(m.Chat.ID, "Keyingi buyruqni tanlang: ")
	msg.ReplyToMessageID = m.MessageID
	msg.ReplyMarkup = buttons.Todo
	h.bot.Send(msg)
}

func (h *Handler) AddTodo(m *tg.Message) {
	text := word.TODO
	msg := tg.NewMessage(m.Chat.ID, text)
	msg.ReplyMarkup = buttons.New_todo
	msg.ReplyToMessageID = m.MessageID
	h.bot.Send(msg)
}

func (h *Handler) AddTitle(m *tg.Message) {

	if err := h.srvc.UserService.SetAction(m.Chat.ID, action.TODO_NEW_TITLE); err != nil {
		fmt.Println(err.Error() + " in addtitle")
	}
	msg := tg.NewMessage(m.Chat.ID, "Titleni kiriting : ")
	msg.ReplyToMessageID = m.MessageID
	msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
	h.bot.Send(msg)
}

func (h *Handler) AddDescription(m *tg.Message) {

	if err := h.srvc.UserService.SetAction(m.Chat.ID, action.TODO_NEW_DESCRIPTION); err != nil {
		fmt.Println(err.Error() + " in adddescription")
	}
	msg := tg.NewMessage(m.Chat.ID, "Description kiriting : ")
	msg.ReplyToMessageID = m.MessageID
	msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
	h.bot.Send(msg)
}

// func (h *Handler) AddPhoto(m *tg.Message) {

// 	h.rds.Set(fmt.Sprint(m.Chat.ID), action.TODO_NEW_PICTURE)
// 	msg := tg.NewMessage(m.Chat.ID, "Send Photo : ")
// 	msg.ReplyToMessageID = m.MessageID
// 	h.bot.Send(msg)
// }

// func (h *Handler) AddFile(m *tg.Message) {

// 	h.rds.Set(fmt.Sprint(m.Chat.ID), action.TODO_NEW_FILE)
// 	msg := tg.NewMessage(m.Chat.ID, "Send File : ")
// 	msg.ReplyToMessageID = m.MessageID
// 	h.bot.Send(msg)
// }

func (h *Handler) AddNotification(m *tg.Message) {

	h.srvc.UserService.SetAction(m.Chat.ID, action.TODO_NEW_NOTIFICATION)
	msg := tg.NewMessage(m.Chat.ID, "Vaqtni tanlang : ")
	msg.ReplyToMessageID = m.MessageID
	msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
	msg.ReplyMarkup = buttons.Clock
	h.bot.Send(msg)
}
