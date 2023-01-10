package updates

import (
	"fmt"
	"strings"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/buttons"
	"github.com/realtemirov/projects/tgbot/helper"
	"github.com/realtemirov/projects/tgbot/helper/const/action"
	"github.com/realtemirov/projects/tgbot/helper/const/word"

	"github.com/realtemirov/projects/tgbot/service"
)

type Handler struct {
	srvc service.Service
	bot  *tg.BotAPI
}

func NewHandler(s service.Service, b *tg.BotAPI) *Handler {
	return &Handler{
		srvc: s,
		bot:  b,
	}
}

func Message(h *Handler, update *tg.Update) {

	m := update.Message

	act, err := h.srvc.UserService.GetAction(m.Chat.ID)

	if err != nil {
		fmt.Println(err.Error())
		act = action.EMPTY
	}

	switch m.Text {
	case "/me", word.MENU_PROFILE:
		h.Profile(m)
	case "/start":
		h.SingUp(m)
	case "/todo", word.MENU_TODO:
		h.AddTodo(m)
	case "/history", word.TODO_HISTORY:
		if !h.GetAllTodosByUserID(m, true, "history") {
			msg := tg.NewMessage(m.Chat.ID, "You don't have any tasks")
			msg.ReplyMarkup = buttons.Menu
			h.bot.Send(msg)
		}
	case "/list", word.TODO_TODOS:

		if !h.GetAllTodosByUserID(m, false, "list") {
			msg := tg.NewMessage(m.Chat.ID, "You don't have any tasks")
			msg.ReplyMarkup = buttons.Menu
			h.bot.Send(msg)
		}
	case word.TODO_CANCEL:
		h.CancelTodo(m)
	case word.TODO_NEW_SAVE:
		h.DeadlineOrNotification(m)
	default:
		switch act {
		case action.TODO_NEW_TITLE:
			h.SetTodoText(m)
		}
	}
}

func CallbackQuery(h *Handler, update *tg.Update) {

	c := update.CallbackQuery
	m := c.Message
	text := ""
	act, err := h.srvc.UserService.GetAction(m.Chat.ID)
	if err != nil {
		fmt.Println(err.Error())
		act = action.EMPTY
	}

	if act == action.EMPTY {
		data := strings.Split(c.Data, "!")
		done := false
		if data[2] == "history" {
			done = true
		}
		switch data[0] {
		case "id":
			{
				if done {
					h.TodoDone(data[1], data[2], c.Message)
				} else {
					h.TodoView(data[1], data[2], c.Message)
				}
				msg := tg.NewDeleteMessage(c.From.ID, c.Message.MessageID)
				h.bot.Send(msg)
			}
		case "delete":
			{
				if err := h.srvc.TodoService.DeleteByID(data[1]); err != nil {
					fmt.Println(err.Error())
				}

				del := tg.NewDeleteMessage(c.From.ID, c.Message.MessageID)
				h.bot.Send(del)
				if !h.GetAllTodosByUserID(c.Message, done, data[2]) {
					msg := tg.NewMessage(m.Chat.ID, "You don't have any tasks")
					msg.ReplyMarkup = buttons.Menu
					h.bot.Send(msg)
				}
			}
		case "done":
			{
				if err := h.srvc.TodoService.Done(data[1], true); err != nil {
					fmt.Println(err.Error())
				}
				del := tg.NewDeleteMessage(c.From.ID, c.Message.MessageID)
				h.bot.Send(del)
				if !h.GetAllTodosByUserID(c.Message, done, data[2]) {
					msg := tg.NewMessage(m.Chat.ID, "You don't have any tasks")
					msg.ReplyMarkup = buttons.Menu
					h.bot.Send(msg)
				}
			}
		case "restore":
			{
				if err := h.srvc.TodoService.Done(data[1], false); err != nil {
					fmt.Println(err.Error())
				}
				del := tg.NewDeleteMessage(c.From.ID, c.Message.MessageID)
				h.bot.Send(del)
				if !h.GetAllTodosByUserID(c.Message, done, data[2]) {
					msg := tg.NewMessage(m.Chat.ID, "You don't have any tasks")
					msg.ReplyMarkup = buttons.Menu
					h.bot.Send(msg)
				}

			}
		case "back":
			{
				h.bot.Send(tg.NewDeleteMessage(c.From.ID, c.Message.MessageID))
				tg.NewCallback(c.Data, data[2])
				if data[2] == "list" {
					if !h.GetAllTodosByUserID(c.Message, done, data[2]) {
						msg := tg.NewMessage(m.Chat.ID, "You don't have any tasks")
						msg.ReplyMarkup = buttons.Menu
						h.bot.Send(msg)
					}
				} else if data[2] == "history" {
					if !h.GetAllTodosByUserID(c.Message, done, data[2]) {
						msg := tg.NewMessage(m.Chat.ID, "You don't have any tasks")
						msg.ReplyMarkup = buttons.Menu
						h.bot.Send(msg)
					}
				}
			}
		}
	} else if act == action.TODO_NEW_TIME {
		d := strings.Split(c.Data, "!")
		del := tg.NewDeleteMessage(m.Chat.ID, m.MessageID)
		h.bot.Send(del)
		switch d[0] {
		case "deadline":
			{
				// h.srvc.UserService.SetAction(m.Chat.ID, action.TODO_NEW_DEADLINE)
				msg := tg.NewMessage(m.Chat.ID, "Choose month")
				msg.ReplyMarkup = buttons.Month
				h.bot.Send(msg)
			}
		case "notification":
			{
				// h.srvc.UserService.SetAction(m.Chat.ID, action.TODO_NEW_NOTIFICATION)
				msg := tg.NewMessage(m.Chat.ID, "Choose hour")
				msg.ReplyMarkup = buttons.Hour("notification")
				h.bot.Send(msg)
			}
		case "month":
			{
				t := time.Time{}
				month := d[1]
				switch month {
				case "1":
					t = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "2":
					t = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "3":
					t = time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "4":
					t = time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "5":
					t = time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "6":
					t = time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "7":
					t = time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "8":
					t = time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "9":
					t = time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "10":
					t = time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "11":
					t = time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "12":
					t = time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				}
				msg := tg.NewMessage(m.Chat.ID, "Choose day")
				var btn tg.InlineKeyboardMarkup
				if month == "2" {
					btn = buttons.Day(28)
				} else if month == "1" || month == "3" || month == "5" || month == "7" || month == "8" || month == "10" || month == "12" {
					btn = buttons.Day(31)
				} else if month == "4" || month == "6" || month == "9" || month == "11" {
					btn = buttons.Day(30)
				}
				msg.ReplyMarkup = btn
				h.bot.Send(msg)

			}
		case "day":
			{
				day := d[1]
				t := time.Time{}
				switch day {
				case "1":
					t = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "2":
					t = time.Date(1, 1, 2, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "3":
					t = time.Date(1, 1, 3, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "4":
					t = time.Date(1, 1, 4, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "5":
					t = time.Date(1, 1, 5, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "6":
					t = time.Date(1, 1, 6, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "7":
					t = time.Date(1, 1, 7, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "8":
					t = time.Date(1, 1, 8, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "9":
					t = time.Date(1, 1, 9, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "10":
					t = time.Date(1, 1, 10, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "11":
					t = time.Date(1, 1, 11, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "12":
					t = time.Date(1, 1, 12, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "13":
					t = time.Date(1, 1, 13, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "14":
					t = time.Date(1, 1, 14, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "15":
					t = time.Date(1, 1, 15, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "16":
					t = time.Date(1, 1, 16, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "17":
					t = time.Date(1, 1, 17, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "18":
					t = time.Date(1, 1, 18, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "19":
					t = time.Date(1, 1, 19, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "20":
					t = time.Date(1, 1, 20, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "21":
					t = time.Date(1, 1, 21, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "22":
					t = time.Date(1, 1, 22, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "23":
					t = time.Date(1, 1, 23, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "24":
					t = time.Date(1, 1, 24, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "25":
					t = time.Date(1, 1, 25, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "26":
					t = time.Date(1, 1, 26, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "27":
					t = time.Date(1, 1, 27, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "28":
					t = time.Date(1, 1, 28, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "29":
					t = time.Date(1, 1, 29, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "30":
					t = time.Date(1, 1, 30, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				case "31":
					t = time.Date(1, 1, 31, 0, 0, 0, 0, time.UTC)
					h.srvc.TodoService.AddTime(m.Chat.ID, &t)
				}
				msg := tg.NewMessage(m.Chat.ID, "Choose hour")
				msg.ReplyMarkup = buttons.Hour("deadline")
				h.bot.Send(msg)
			}
		case "clock":
			{

				t := helper.Clock_Hour(d[1])

				todo, err := h.srvc.TodoService.AddHour(m.Chat.ID, &t, d[2])
				if err != nil {
					fmt.Println(err.Error())
					text = "Error1 : " + err.Error() + "\n"
				}
				err = h.srvc.TodoService.SetIsSet(m.Chat.ID)
				if err != nil {
					fmt.Println(err.Error())
					text += "Error2 : " + err.Error() + "\n"
				}
				err = h.srvc.UserService.SetAction(m.Chat.ID, action.EMPTY)
				if err != nil {
					fmt.Println(err.Error())
					text += "Error3 : " + err.Error() + "\n"
				}
				if text == "" {
					text = todo.ToString()
				}

				msg := tg.NewMessage(m.Chat.ID, text)
				msg.ReplyMarkup = buttons.Menu
				msg.ParseMode = "HTML"
				h.bot.Send(msg)
			}

		}

	}
}
