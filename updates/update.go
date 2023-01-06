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
	case "/todo", word.TODO_ADD:
		h.AddTodo(m)
	case "/history", word.TODO_HISTORY:
		// todo
	case "/list", word.TODO_TODOS:
		// todo
	case "/menu", word.MENU_TODO:
		h.Todo(m)
	case word.TODO_NEW_CANCEL:
		h.CancelNew(m)
	default:
		switch act {
		case action.TODO_NEW_TITLE:
			h.SetTodoTitle(m)
		case action.TODO_NEW_DESCRIPTION:
			h.SetTodoDescription(m)
		case action.TODO_NEW_PICTURE:
			//h.SetTodoPhoto(m)
		case action.TODO_NEW_FILE:
			//h.SetTodoFile(m)
		case action.TODO_ADD:
			{
				switch m.Text {
				case word.TODO_NEW_DESCRIPTION:
					h.AddDescription(m)
				case word.TODO_NEW_PICTURE:
				//	h.AddPhoto(m)
				case word.TODO_NEW_FILE:
					//h.AddFile(m)
				case word.TODO_NEW_SAVE:
					h.AddNotification(m)
				default:

				}
			}
		case action.EMPTY:
			{
				switch m.Text {
				case word.TODO_NEW_TITLE:
					h.AddTitle(m)
				}

			}
		}
	}
}

func CallbackQuery(h *Handler, update *tg.Update) {

	fmt.Println("CallbackQuery: ", update.CallbackQuery.Data, "from", update.CallbackQuery.From.ID)
	c := update.CallbackQuery
	m := c.Message
	text := ""
	act, err := h.srvc.UserService.GetAction(m.Chat.ID)
	if err != nil {
		fmt.Println(err.Error())
		act = action.EMPTY
	}
	if act == action.TODO_NEW_NOTIFICATION {
		if strings.HasPrefix(c.Data, "clock") {

			del := tg.NewDeleteMessage(m.Chat.ID, m.MessageID)
			h.bot.Send(del)

			t := time.Time{}.Add(helper.Clock_Hour(c.Data))

			err = h.srvc.TodoService.AddNotification(m.Chat.ID, &t)
			if err != nil {
				fmt.Println(err.Error())
				text = "Error : " + err.Error() + "\n"
			}
			err = h.srvc.TodoService.SetIsSet(m.Chat.ID)
			if err != nil {
				fmt.Println(err.Error())
				text += "Error : " + err.Error() + "\n"
			}
			err = h.srvc.UserService.SetAction(m.Chat.ID, action.EMPTY)
			if err != nil {
				fmt.Println(err.Error())
				text += "Error : " + err.Error() + "\n"
			}
		}
	}

	if text == "" {
		text = "Done"
	}

	msg := tg.NewMessage(m.Chat.ID, text)
	msg.ReplyMarkup = buttons.Menu
	h.bot.Send(msg)

}
