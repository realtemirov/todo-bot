package updates

import (
	"fmt"
	"strconv"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/buttons"
	"github.com/realtemirov/projects/tgbot/helper/const/action"
	"github.com/realtemirov/projects/tgbot/model"
)

/*
func CallbackQuery(bot *tg.BotAPI, update *tg.Update) {
	c := update.CallbackQuery
	switch c.Data {
	// todo buttons
	case query.TODO_ADD:
		edt := tg.NewEditMessageReplyMarkup(c.Message.Chat.ID, c.Message.MessageID, buttons.New_todo)
		bot.Send(edt)
	default:
		Alert(bot, update, c.Data)
	}

}*/

func (h *Handler) TodoAdd(c *tg.CallbackQuery) {
	msg := tg.NewMessage(c.Message.Chat.ID, "Your action is "+c.Data)
	msg.ReplyMarkup = buttons.New_todo
	h.bot.Send(msg)
}

func (h *Handler) TodoHistory(c *tg.CallbackQuery) {
	msg := tg.NewMessage(c.Message.Chat.ID, "Your action is "+c.Data)
	h.bot.Send(msg)
	Alert(h.bot, c)
}

func (h *Handler) Todos(c *tg.CallbackQuery) {
	msg := tg.NewMessage(c.Message.Chat.ID, "Your action is "+c.Data)
	h.bot.Send(msg)
	Alert(h.bot, c)
}

func (h *Handler) TodoNewTitle(c *tg.CallbackQuery) {
	m := c.Message
	err := h.rds.Set(fmt.Sprint(m.Chat.ID), action.TODO_NEW_TITLE)
	if err != nil {
		fmt.Println(err.Error())
	}
	msg := tg.NewMessage(m.Chat.ID, "Enter title")
	h.bot.Send(msg)
}

func (h *Handler) TodoNewDescription(c *tg.CallbackQuery) {
	m := c.Message
	err := h.rds.Set(fmt.Sprint(m.Chat.ID), action.TODO_NEW_DESCRIPTION)
	if err != nil {
		fmt.Println(err.Error())
	}
	msg := tg.NewMessage(m.Chat.ID, "Enter Description")
	h.bot.Send(msg)
}

func Alert(bot *tg.BotAPI, c *tg.CallbackQuery) {
	resp, err := bot.Request(tg.CallbackConfig{
		CallbackQueryID: c.ID,
		Text:            c.Data,
		ShowAlert:       true,
		CacheTime:       10,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
}

func (h *Handler) TodoNewNotification(c *tg.CallbackQuery) {
	m := c.Message

	err := h.rds.Set(fmt.Sprint(m.Chat.ID), action.TODO_NEW_NOTIFICATION)
	if err != nil {
		fmt.Println(err.Error())
	}
	msg := tg.NewMessage(m.Chat.ID, "Enter Number notification")
	h.bot.Send(msg)
}

func (h *Handler) TodoSetNotification(c *tg.CallbackQuery) {

	m := c.Message
	key := fmt.Sprint(m.Chat.ID)
	n, err := h.rds.Get("notif-" + key)
	if err != nil {
		msg := tg.NewMessage(m.Chat.ID, err.Error())
		h.bot.Send(msg)
		return
	}

	int_n, err := strconv.Atoi(n)
	if err != nil {
		msg := tg.NewMessage(m.Chat.ID, err.Error())
		h.bot.Send(msg)
		return
	}


	id := h.srvc.NotifService.Create(&model.Notification{
		Todo_ID: ,
	})
	if int_n > 0 {

	}
}
