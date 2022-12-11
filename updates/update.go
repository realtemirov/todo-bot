package updates

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/helper/const/action"
	"github.com/realtemirov/projects/tgbot/helper/const/query"
	"github.com/realtemirov/projects/tgbot/helper/const/word"
	"github.com/realtemirov/projects/tgbot/service"
	"github.com/realtemirov/projects/tgbot/storage/redis"
)

type Handler struct {
	srvc service.Service

	rds *redis.RedisCache
	bot *tg.BotAPI
}

func NewHandler(s service.Service, r *redis.RedisCache, b *tg.BotAPI) *Handler {
	return &Handler{
		srvc: s,
		bot:  b,
		rds:  r,
	}
}

func Message(h *Handler, update *tg.Update) {
	m := update.Message
	if m.Photo != nil {
		fmt.Println(m.Text)
		fmt.Println(m.Photo)
	}
	switch m.Text {
	case "/start":
		h.SingUp(m)
	case word.MENU_TODO:
		h.Todo(m)
	case word.MENU_CHALLANGES:
		fmt.Println("Your action is " + word.MENU_CHALLANGES + ", Soon")
	case word.MENU_SETTINGS:
		fmt.Println("Your action is " + word.MENU_SETTINGS + ", Soon")
	case word.MENU_RECOMMENDATION:
		fmt.Println("Your action is " + word.MENU_RECOMMENDATION + ", Soon")
	default:

		act, err := Action(h, fmt.Sprint(m.Chat.ID))

		if err != nil {
			fmt.Println(err.Error())
		}

		text := "Your action is " + act

		switch act {
		case action.TODO_ADD:
			Warning(h, m)
		case action.TODO_NEW_TITLE:
			if !SetTodoTitle(h, m) {
				Warning(h, m)
			}
		case action.TODO_NEW_DESCRIPTION:
			fmt.Println("TODO_CREATE_DESCRIPTION")
		case action.TODO_NEW_DEADLINE:
			fmt.Println("TODO_CREATE_DEADLINE")
		case action.TODO_NEW_PICTURE:
			fmt.Println("TODO_CREATE_PICTURE")
		case action.TODO_NEW_FILE:
			fmt.Println("TODO_CREATE_FILE")
		case action.TODO_NEW_SAVE:
			fmt.Println("TODO_CREATE_SAVE")
		case action.TODO_NEW_CANCEL:
			fmt.Println("TODO_CREATE_CANCEL")
		case action.EMPTY:
			fmt.Println("EMPTY")
		default:
			fmt.Println("default")
		}

		msg := tg.NewMessage(update.Message.Chat.ID, text)
		h.bot.Send(msg)
	}
}

func CallbackQuery(h *Handler, update *tg.Update) {

	c := update.CallbackQuery

	act, err := Action(h, fmt.Sprint(c.Message.Chat.ID))

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("here1")
		msg := tg.NewMessage(c.Message.Chat.ID, "Something is wrong")
		h.bot.Send(msg)
		return
	}

	if act == action.TODO_ADD && c.Data != query.TODO_NEW_SAVE && c.Data != query.TODO_ADD {
		fmt.Println(act)
		fmt.Println("here2")
		msg := tg.NewMessage(c.Message.Chat.ID, "Please save your last todo")
		h.bot.Send(msg)
		return
	}

	switch c.Data {
	// todo buttons
	case query.TODO_ADD:
		h.TodoAdd(c)
	case query.TODO_HISTORY:
		h.TodoHistory(c)
	case query.TODO_TODOS:
		h.Todos(c)
	case query.TODO_NEW_TITLE:
		h.TodoNewTitle(c)
	case query.TODO_NEW_DESCRIPTION:
		fmt.Println("TODO_NEW_DESCRIPTION")
	case query.TODO_NEW_DEADLINE:
		fmt.Println("TODO_NEW_DEADLINE")
	case query.TODO_NEW_NOTIFICATON:
		fmt.Println("TODO_NEW_NOTIFICATON")
	case query.TODO_NEW_PICTURE:
		fmt.Println("TODO_NEW_PICTURE")
	case query.TODO_NEW_FILE:
		fmt.Println("TODO_NEW_FILE")
	case query.TODO_NEW_SAVE:
		Notification(h, c.Message)
	case query.TODO_NEW_CANCEL:
		fmt.Println("TODO_NEW_CANCEL")
	case query.NOTIFICATION_ON:
		fmt.Println("NOTIFICATION_ON")
		SaveTodo(h, c.Message)
	case query.NOTIFICATION_OFF:
		fmt.Println("NOTIFICATION_OFF")
	default:
		Alert(h.bot, c)
	}

}
func Action(h *Handler, key string) (string, error) {

	act, err := h.rds.Get(key)
	fmt.Println("act", act)
	if err != nil {
		return "empty", err
	}
	return act, nil
}
