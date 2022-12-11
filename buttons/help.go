package buttons

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/helper/const/query"
	"github.com/realtemirov/projects/tgbot/helper/const/word"
)

var Notification = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.NOTIFICATION_ON, query.NOTIFICATION_ON),
		//list of tasks
		tg.NewInlineKeyboardButtonData(word.NOTIFICATION_OFF, query.NOTIFICATION_OFF),
	),
)

var Clock = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.CLOCK_1, query.CLOCK_1),
		tg.NewInlineKeyboardButtonData(word.CLOCK_2, query.CLOCK_2),
		tg.NewInlineKeyboardButtonData(word.CLOCK_3, query.CLOCK_3),
		tg.NewInlineKeyboardButtonData(word.CLOCK_4, query.CLOCK_4),
		tg.NewInlineKeyboardButtonData(word.CLOCK_5, query.CLOCK_5),
		tg.NewInlineKeyboardButtonData(word.CLOCK_6, query.CLOCK_6),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.CLOCK_7, query.CLOCK_7),
		tg.NewInlineKeyboardButtonData(word.CLOCK_8, query.CLOCK_8),
		tg.NewInlineKeyboardButtonData(word.CLOCK_9, query.CLOCK_9),
		tg.NewInlineKeyboardButtonData(word.CLOCK_10, query.CLOCK_10),
		tg.NewInlineKeyboardButtonData(word.CLOCK_11, query.CLOCK_11),
		tg.NewInlineKeyboardButtonData(word.CLOCK_12, query.CLOCK_12),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.CLOCK_13, query.CLOCK_13),
		tg.NewInlineKeyboardButtonData(word.CLOCK_14, query.CLOCK_14),
		tg.NewInlineKeyboardButtonData(word.CLOCK_15, query.CLOCK_15),
		tg.NewInlineKeyboardButtonData(word.CLOCK_16, query.CLOCK_16),
		tg.NewInlineKeyboardButtonData(word.CLOCK_17, query.CLOCK_17),
		tg.NewInlineKeyboardButtonData(word.CLOCK_18, query.CLOCK_18),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.CLOCK_19, query.CLOCK_19),
		tg.NewInlineKeyboardButtonData(word.CLOCK_20, query.CLOCK_20),
		tg.NewInlineKeyboardButtonData(word.CLOCK_21, query.CLOCK_21),
		tg.NewInlineKeyboardButtonData(word.CLOCK_22, query.CLOCK_22),
		tg.NewInlineKeyboardButtonData(word.CLOCK_23, query.CLOCK_23),
		tg.NewInlineKeyboardButtonData(word.CLOCK_24, query.CLOCK_24),
	),
)
