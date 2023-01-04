package updates

import (
	"fmt"
	"io/ioutil"

	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/buttons"
	"github.com/realtemirov/projects/tgbot/helper/const/action"
	"github.com/realtemirov/projects/tgbot/model"
)

func (h *Handler) Getme(m *tg.Message) {

	user, err := h.srvc.UserService.GetByID(m.Chat.ID)
	if err != nil {
		fmt.Println(err.Error() + " in getme")
	}
	fmt.Println(user)
	var p tg.PhotoConfig
	if user == nil || strings.Contains(user.Photo_URL, "image.jpg") {
		fmt.Println("user is null or photo is null")
		photoBytes, err := ioutil.ReadFile("image.jpg")
		if err != nil {
			fmt.Println(err.Error())
		}
		photoFileBytes := tg.FileBytes{Name: "picture", Bytes: photoBytes}
		p = tg.NewPhoto(user.ID, photoFileBytes)
	} else {
		p = tg.NewPhoto(user.ID, tg.FileID(user.Photo_URL))
	}
	p.Caption = fmt.Sprintf("<b>👤 Fullname : <i>%s %s</i> \n🌐 Username : %s \n🆔 User ID : <code>%d</code>\n📝 Sign Up : <code>%v</code></b>", user.FirsName, user.LastName, "@"+user.Username, user.ID, user.CreatedAt)
	p.ParseMode = "HTML"
	h.bot.Send(p)

}

func (h *Handler) SingUp(m *tg.Message) {
	user := &model.User{
		ID:        m.Chat.ID,
		Username:  m.Chat.UserName,
		FirsName:  m.Chat.FirstName,
		LastName:  m.Chat.LastName,
		IsAdmin:   false,
		Photo_URL: "image.jpg",
	}

	cnf := tg.UserProfilePhotosConfig{
		UserID: m.Chat.ID,
		Offset: 0,
		Limit:  1,
	}

	photos, err := h.bot.GetUserProfilePhotos(cnf)

	if err == nil && photos.TotalCount != 0 {
		user.Photo_URL = photos.Photos[0][0].FileID
	}

	id, err := h.srvc.UserService.Add(user)
	text := ""
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			text = fmt.Sprintf("%s siz allaqachon ro'yxatdan o'tgansiz.", m.Chat.FirstName)
		} else {
			text = fmt.Sprintf("Error %s", err.Error())
		}

	} else {
		text = fmt.Sprintf("Salom, *%s*, sizning ID : `%v` \nKeyingi amalni tanlang: ", m.Chat.FirstName, id)
	}
	h.rds.Set(fmt.Sprintf("%v", m.Chat.ID), action.EMPTY)
	msg := tg.NewMessage(m.Chat.ID, text)
	msg.ReplyToMessageID = m.MessageID
	msg.ParseMode = "markdown"
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
	text := `
		Title - Bu taskni nomi.
	Description - Taskni to'liq bayon qilishingiz mumkin.
	Deadline
	`
	msg := tg.NewMessage(m.Chat.ID, text)
	msg.ReplyMarkup = buttons.New_todo
	msg.ReplyToMessageID = m.MessageID
	h.bot.Send(msg)
}

func (h *Handler) AddTitle(m *tg.Message) {

	h.rds.Set(fmt.Sprint(m.Chat.ID), action.TODO_NEW_TITLE)
	msg := tg.NewMessage(m.Chat.ID, "Titleni kiriting : ")
	msg.ReplyToMessageID = m.MessageID
	h.bot.Send(msg)
}

func (h *Handler) AddDescription(m *tg.Message) {

	h.rds.Set(fmt.Sprint(m.Chat.ID), action.TODO_NEW_DESCRIPTION)
	msg := tg.NewMessage(m.Chat.ID, "Description kiriting : ")
	msg.ReplyToMessageID = m.MessageID
	h.bot.Send(msg)
}

func (h *Handler) AddPhoto(m *tg.Message) {

	h.rds.Set(fmt.Sprint(m.Chat.ID), action.TODO_NEW_PICTURE)
	msg := tg.NewMessage(m.Chat.ID, "Send Photo : ")
	msg.ReplyToMessageID = m.MessageID
	h.bot.Send(msg)
}

func (h *Handler) AddFile(m *tg.Message) {

	h.rds.Set(fmt.Sprint(m.Chat.ID), action.TODO_NEW_FILE)
	msg := tg.NewMessage(m.Chat.ID, "Send File : ")
	msg.ReplyToMessageID = m.MessageID
	h.bot.Send(msg)
}

func (h *Handler) AddNotification(m *tg.Message) {

	h.rds.Set(fmt.Sprint(m.Chat.ID), action.TODO_NEW_NOTIFICATION)
	msg := tg.NewMessage(m.Chat.ID, "Vaqtni tanlang : ")
	msg.ReplyToMessageID = m.MessageID
	msg.ReplyMarkup = buttons.Clock
	h.bot.Send(msg)
}
