package helper

import (
	"fmt"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/helper/const/action"
	"github.com/realtemirov/projects/tgbot/model"
)

func ContainsStrings(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Clock_Hour(t string) time.Duration {
	var t2 time.Duration
	switch t {
	case action.CLOCK_1:
		t2 = time.Hour
	case action.CLOCK_2:
		t2 = time.Hour * 2
	case action.CLOCK_3:
		t2 = time.Hour * 3
	case action.CLOCK_4:
		t2 = time.Hour * 4
	case action.CLOCK_5:
		t2 = time.Hour * 5
	case action.CLOCK_6:
		t2 = time.Hour * 6
	case action.CLOCK_7:
		t2 = time.Hour * 7
	case action.CLOCK_8:
		t2 = time.Hour * 8
	case action.CLOCK_9:
		t2 = time.Hour * 9
	case action.CLOCK_10:
		t2 = time.Hour * 10
	case action.CLOCK_11:
		t2 = time.Hour * 11
	case action.CLOCK_12:
		t2 = time.Hour * 12
	case action.CLOCK_13:
		t2 = time.Hour * 13
	case action.CLOCK_14:
		t2 = time.Hour * 14
	case action.CLOCK_15:
		t2 = time.Hour * 15
	case action.CLOCK_16:
		t2 = time.Hour * 16
	case action.CLOCK_17:
		t2 = time.Hour * 17
	case action.CLOCK_18:
		t2 = time.Hour * 18
	case action.CLOCK_19:
		t2 = time.Hour * 19
	case action.CLOCK_20:
		t2 = time.Hour * 20
	case action.CLOCK_21:
		t2 = time.Hour * 21
	case action.CLOCK_22:
		t2 = time.Hour * 22
	case action.CLOCK_23:
		t2 = time.Hour * 23
	case action.CLOCK_24:
		t2 = time.Hour * 24
	}
	return t2
}

func UserToString(user *tg.User, user2 *model.User) string {
	text := fmt.Sprintf("<b>🆔 ID: <code>%d</code> \n👤 Fullname: %s %s</b>\n", user.ID, user.FirstName, user.LastName)
	if user.UserName != "" {
		text += fmt.Sprintf("<b>🌐 Username: %s</b>\n", "@"+user.UserName)
	}
	if user.LanguageCode != "" {
		text += fmt.Sprintf("<b>🌐 Language: %s</b>\n", user.LanguageCode)
	}

	text += fmt.Sprintf("<b>⏳ Created at: <code>%s</code></b>\n", user2.CreatedAt.Format("01 January 2006 15:04:05"))
	return text
}
