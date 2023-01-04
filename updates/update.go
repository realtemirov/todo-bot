package updates

import (
	"fmt"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/realtemirov/projects/tgbot/helper"
	"github.com/realtemirov/projects/tgbot/helper/const/action"
	"github.com/realtemirov/projects/tgbot/helper/const/query"
	"github.com/realtemirov/projects/tgbot/helper/const/word"
	"github.com/realtemirov/projects/tgbot/helper/convert"
	"github.com/realtemirov/projects/tgbot/model"
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

	act, err := h.rds.Get(fmt.Sprint(m.Chat.ID))

	if err != nil {
		sendError(err, m, h)
	}

	switch m.Text {
	case "/me", word.MENU_PROFILE:
		h.Getme(m)
	case "/start":
		h.SingUp(m)
	case "/todo", word.TODO_ADD:
		h.AddTodo(m)
	case word.TODO_HISTORY:
		// todo
	case word.TODO_TODOS:
		// todo
	case word.MENU_TODO:
		h.Todo(m)
	case word.MENU_CHALLANGES:
		// todo
	case word.MENU_RECOMMENDATION:
		// todo
	case word.MENU_SETTINGS:
		// todo
	case word.TODO_NEW_CANCEL:
		h.CancelNew(m)
		// todo
	case word.TODO_NEW_TITLE:
		if act == action.EMPTY {
			h.AddTitle(m)
		} else {
			Please(h, m)
		}
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
					h.AddPhoto(m)
				case word.TODO_NEW_FILE:
					h.AddFile(m)
				case word.TODO_NEW_SAVE:
					h.AddNotification(m)
				default:
					Please(h, m)
				}
			}
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

	id := uuid.New().String()
	if act == action.TODO_NEW_NOTIFICATION {
		h.SaveTodo(update.CallbackQuery.Message, id)
		//t := time.Time{}
		n := &model.Notification{
			Base: model.Base{
				ID: id,
			},
			Todo_ID:    id,
			User_ID:    c.Message.Chat.ID,
			Notif_date: time.Now().Add(time.Second * 5),
		}
		n_id, err := h.srvc.NotifService.Create(n)
		if err != nil {
			fmt.Println(err.Error())
		} else if n_id != id {
			fmt.Printf("New id : %s,   last id : %s", n_id, id)
		}

	}
	// todo
	fmt.Println("hello", update.CallbackQuery.Data)
}

func Action(h *Handler, key string) (string, error) {

	act, err := h.rds.Get(key)

	if err != nil {
		return "", err
	}
	return act, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func sendError(err error, m *tg.Message, h *Handler) {
	msg := tg.NewMessage(m.Chat.ID, "Something is wrong : "+err.Error())
	h.bot.Send(msg)
}

func Please(h *Handler, m *tg.Message) {
	json, err := h.rds.Get("todo-" + fmt.Sprint(m.Chat.ID))
	if err != nil {

		sendError(err, m, h)
	}
	todo, err := convert.StringToTodo(json)
	if err != nil {
		fmt.Println("here")
		sendError(err, m, h)
	}
	text := helper.TodoToString(todo)
	msg := tg.NewMessage(m.Chat.ID, "Please save last todo\n"+text)
	msg.ReplyToMessageID = m.MessageID

	msg.ParseMode = "markdown"
	h.bot.Send(msg)
}

func notif_clock(t string) time.Duration {
	var t2 time.Duration
	switch t {
	case query.CLOCK_1:
		t2 = time.Hour
	case query.CLOCK_2:
		t2 = time.Hour * 2
	case query.CLOCK_3:
		t2 = time.Hour * 3
	case query.CLOCK_4:
		t2 = time.Hour * 4
	case query.CLOCK_5:
		t2 = time.Hour * 5
	case query.CLOCK_6:
		t2 = time.Hour * 6
	case query.CLOCK_7:
		t2 = time.Hour * 7
	case query.CLOCK_8:
		t2 = time.Hour * 8
	case query.CLOCK_9:
		t2 = time.Hour * 9
	case query.CLOCK_10:
		t2 = time.Hour * 10
	case query.CLOCK_11:
		t2 = time.Hour * 11
	case query.CLOCK_12:
		t2 = time.Hour * 12
	case query.CLOCK_13:
		t2 = time.Hour * 13
	case query.CLOCK_14:
		t2 = time.Hour * 14
	case query.CLOCK_15:
		t2 = time.Hour * 15
	case query.CLOCK_16:
		t2 = time.Hour * 16
	case query.CLOCK_17:
		t2 = time.Hour * 17
	case query.CLOCK_18:
		t2 = time.Hour * 18
	case query.CLOCK_19:
		t2 = time.Hour * 19
	case query.CLOCK_20:
		t2 = time.Hour * 20
	case query.CLOCK_21:
		t2 = time.Hour * 21
	case query.CLOCK_22:
		t2 = time.Hour * 22
	case query.CLOCK_23:
		t2 = time.Hour * 23
	case query.CLOCK_24:
		t2 = time.Hour * 24
	}
	return t2
}
