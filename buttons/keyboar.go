package buttons

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/helper/const/word"
)

var Menu = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(word.MENU_TODO),
		tg.NewKeyboardButton(word.MENU_PROFILE),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(word.MENU_SETTINGS),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(word.MENU_RECOMMENDATION),
	),
)
