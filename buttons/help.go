package buttons

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/helper/const/action"
	"github.com/realtemirov/projects/tgbot/helper/const/word"
)

var Clock = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.CLOCK_1, action.CLOCK_1),
		tg.NewInlineKeyboardButtonData(word.CLOCK_2, action.CLOCK_2),
		tg.NewInlineKeyboardButtonData(word.CLOCK_3, action.CLOCK_3),
		tg.NewInlineKeyboardButtonData(word.CLOCK_4, action.CLOCK_4),
		tg.NewInlineKeyboardButtonData(word.CLOCK_5, action.CLOCK_5),
		tg.NewInlineKeyboardButtonData(word.CLOCK_6, action.CLOCK_6),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.CLOCK_7, action.CLOCK_7),
		tg.NewInlineKeyboardButtonData(word.CLOCK_8, action.CLOCK_8),
		tg.NewInlineKeyboardButtonData(word.CLOCK_9, action.CLOCK_9),
		tg.NewInlineKeyboardButtonData(word.CLOCK_10, action.CLOCK_10),
		tg.NewInlineKeyboardButtonData(word.CLOCK_11, action.CLOCK_11),
		tg.NewInlineKeyboardButtonData(word.CLOCK_12, action.CLOCK_12),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.CLOCK_13, action.CLOCK_13),
		tg.NewInlineKeyboardButtonData(word.CLOCK_14, action.CLOCK_14),
		tg.NewInlineKeyboardButtonData(word.CLOCK_15, action.CLOCK_15),
		tg.NewInlineKeyboardButtonData(word.CLOCK_16, action.CLOCK_16),
		tg.NewInlineKeyboardButtonData(word.CLOCK_17, action.CLOCK_17),
		tg.NewInlineKeyboardButtonData(word.CLOCK_18, action.CLOCK_18),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.CLOCK_19, action.CLOCK_19),
		tg.NewInlineKeyboardButtonData(word.CLOCK_20, action.CLOCK_20),
		tg.NewInlineKeyboardButtonData(word.CLOCK_21, action.CLOCK_21),
		tg.NewInlineKeyboardButtonData(word.CLOCK_22, action.CLOCK_22),
		tg.NewInlineKeyboardButtonData(word.CLOCK_23, action.CLOCK_23),
		tg.NewInlineKeyboardButtonData(word.CLOCK_24, action.CLOCK_24),
	),
)

var Day = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(action.DAY_1, action.DAY_1),
		tg.NewInlineKeyboardButtonData(action.DAY_2, action.DAY_2),
		tg.NewInlineKeyboardButtonData(action.DAY_3, action.DAY_3),
		tg.NewInlineKeyboardButtonData(action.DAY_4, action.DAY_4),
		tg.NewInlineKeyboardButtonData(action.DAY_5, action.DAY_5),
		tg.NewInlineKeyboardButtonData(action.DAY_6, action.DAY_6),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(action.DAY_7, action.DAY_7),
		tg.NewInlineKeyboardButtonData(action.DAY_8, action.DAY_8),
		tg.NewInlineKeyboardButtonData(action.DAY_9, action.DAY_9),
		tg.NewInlineKeyboardButtonData(action.DAY_10, action.DAY_10),
		tg.NewInlineKeyboardButtonData(action.DAY_11, action.DAY_11),
		tg.NewInlineKeyboardButtonData(action.DAY_12, action.DAY_12),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(action.DAY_13, action.DAY_13),
		tg.NewInlineKeyboardButtonData(action.DAY_14, action.DAY_14),
		tg.NewInlineKeyboardButtonData(action.DAY_15, action.DAY_15),
		tg.NewInlineKeyboardButtonData(action.DAY_16, action.DAY_16),
		tg.NewInlineKeyboardButtonData(action.DAY_17, action.DAY_17),
		tg.NewInlineKeyboardButtonData(action.DAY_18, action.DAY_18),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(action.DAY_19, action.DAY_19),
		tg.NewInlineKeyboardButtonData(action.DAY_20, action.DAY_20),
		tg.NewInlineKeyboardButtonData(action.DAY_21, action.DAY_21),
		tg.NewInlineKeyboardButtonData(action.DAY_22, action.DAY_22),
		tg.NewInlineKeyboardButtonData(action.DAY_23, action.DAY_23),
		tg.NewInlineKeyboardButtonData(action.DAY_24, action.DAY_24),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(action.DAY_25, action.DAY_25),
		tg.NewInlineKeyboardButtonData(action.DAY_26, action.DAY_26),
		tg.NewInlineKeyboardButtonData(action.DAY_27, action.DAY_27),
		tg.NewInlineKeyboardButtonData(action.DAY_28, action.DAY_28),
		tg.NewInlineKeyboardButtonData(action.DAY_29, action.DAY_29),
		tg.NewInlineKeyboardButtonData(action.DAY_30, action.DAY_30),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(action.DAY_31, action.DAY_31),
	),
)

var Ok = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(word.TODO_OK),
		tg.NewKeyboardButton(word.TODO_CANCEL),
	),
)
