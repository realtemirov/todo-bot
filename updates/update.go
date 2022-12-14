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

var (
	actions = []string{
		action.TODO_ADD,
		action.TODO_HISTORY,
		action.TODO_TODOS,
		action.TODO_NEW_SAVE,
		action.TODO_NEW_CANCEL,
		action.EMPTY,
	}

	clocks = []string{
		query.CLOCK_1,
		query.CLOCK_2,
		query.CLOCK_3,
		query.CLOCK_4,
		query.CLOCK_5,
		query.CLOCK_6,
		query.CLOCK_7,
		query.CLOCK_8,
		query.CLOCK_9,
		query.CLOCK_10,
		query.CLOCK_11,
		query.CLOCK_12,
		query.CLOCK_13,
		query.CLOCK_14,
		query.CLOCK_15,
		query.CLOCK_16,
		query.CLOCK_17,
		query.CLOCK_18,
		query.CLOCK_19,
		query.CLOCK_20,
		query.CLOCK_21,
		query.CLOCK_22,
		query.CLOCK_23,
		query.CLOCK_24,
	}
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

	switch m.Text {
	case "/start":
		h.SingUp(m)
	case "/me":
		h.Getme(m)
	case word.MENU_TODO:
		h.Todo(m)
	default:
		act, err := Action(h, fmt.Sprint(m.Chat.ID))
		LogErr(err)

		switch act {
		case action.TODO_ADD:
			Warning(h, m)
		case action.TODO_NEW_TITLE:
			if !SetTodoTitle(h, m) {
				Warning(h, m)
			}
		case action.TODO_NEW_DESCRIPTION:
			if !SetTodoDescription(h, m) {
				Warning(h, m)
			}
		case action.TODO_NEW_NOTIFICATION:
			if !SetTodoNotification(h, m) {
				Warning(h, m)
			}
		default:
			delMsg := tg.NewDeleteMessage(m.Chat.ID, m.MessageID)
			h.bot.Send(delMsg)

		}
	}
}

func CallbackQuery(h *Handler, update *tg.Update) {

	c := update.CallbackQuery
	act, err := Action(h, fmt.Sprint(c.Message.Chat.ID))
	fmt.Println(c.Data)
	fmt.Println()
	if err != nil {
		msg := tg.NewMessage(c.Message.Chat.ID, "Something is wrong"+err.Error())
		h.bot.Send(msg)
		return
	}

	if c.Data != act && !contains(actions, act) {

		msg := tg.NewMessage(c.Message.Chat.ID, "Please save your last todo")
		h.bot.Send(msg)
		return
	}

	switch c.Data {

	// MENU
	case query.TODO_ADD:
		h.TodoAdd(c)
	case query.TODO_NEW_TITLE:
		h.TodoNewTitle(c)
	case query.TODO_NEW_DESCRIPTION:
		if act == action.TODO_ADD {
			h.TodoNewDescription(c)
		} else {
			c.Data = "Oldin todo kiriting"
			Alert(h.bot, c)
		}

	case query.TODO_NEW_NOTIFICATION:
		h.TodoNewNotification(c)
	case query.TODO_NEW_SAVE:
		if act == action.TODO_ADD {
			SaveTodo(h, c.Message)
		} else {
			c.Data = "Oldin todo kiriting"
			Alert(h.bot, c)
		}
	case query.TODO_NEW_CANCEL:
		if act == action.TODO_ADD {
			CancelNew(h, c.Message)
		} else {
			c.Data = "Oldin todo kiriting"
			Alert(h.bot, c)
		}
	default:
		if contains(clocks, c.Data) {
			h.TodoSetNotification(c)
		}
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

func LogErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
