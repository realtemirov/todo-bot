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
	if user == nil || strings.Contains(user.Photo_URL, "image.jpg") {
		fmt.Println("user is null or photo is null")
		photoBytes, err := ioutil.ReadFile("image.jpg")
		if err != nil {
			fmt.Println(err.Error())
		}
		photoFileBytes := tg.FileBytes{Name: "picture", Bytes: photoBytes}
		p := tg.NewPhoto(user.ID, photoFileBytes)
		p.Caption = fmt.Sprintf("Hello, *%s*, your ID : `%v`", user.FirsName, user.ID)
		p.ParseMode = "markdown"
		h.bot.Send(p)

	} else {

		h.bot.Send(tg.NewPhoto(user.ID, tg.FileID(user.Photo_URL)))
	}

}

func (h *Handler) SingUp(m *tg.Message) {
	fmt.Println(m.Chat.ID)
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
			text = fmt.Sprintf("You are already registered, %s", m.Chat.FirstName)
		} else {
			text = fmt.Sprintf("Error %s", err.Error())
		}

	} else {
		text = fmt.Sprintf("Hello, *%s*, your ID : `%v` \nChoose these: ", m.Chat.FirstName, id)
	}
	h.rds.Set(fmt.Sprintf("%v", m.Chat.ID), action.EMPTY)
	msg := tg.NewMessage(m.Chat.ID, text)
	msg.ParseMode = "markdown"
	msg.ReplyMarkup = buttons.Menu
	h.bot.Send(msg)
}

func (h *Handler) Todo(m *tg.Message) {
	msg := tg.NewMessage(m.Chat.ID, "Choose these: ")
	msg.ReplyMarkup = buttons.Todo
	h.bot.Send(msg)
}

