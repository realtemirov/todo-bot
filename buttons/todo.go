package buttons

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/helper/const/word"
)

var Todo = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(word.TODO_ADD),
		//list of tasks
		tg.NewKeyboardButton(word.TODO_TODOS),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(word.TODO_HISTORY),
	),
)

var New_todo = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(word.TODO_NEW_TITLE),
		tg.NewKeyboardButton(word.TODO_NEW_DESCRIPTION),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(word.TODO_NEW_PICTURE),
		tg.NewKeyboardButton(word.TODO_NEW_FILE),
	),
	/*	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(word.TODO_NEW_DEADLINE),
		tg.NewKeyboardButton(word.TODO_NEW_NOTIFICATON),
	),*/
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(word.TODO_NEW_SAVE),
		tg.NewKeyboardButton(word.TODO_CANCEL),
	),
)

var Todo_view = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("✅ Done", "view-done-task"),
		tg.NewInlineKeyboardButtonData("🗑 Delete", "view-delete-task"),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("📝 Edit", "view-edit-task"),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("⬅️ ", "view-left-task"),
		tg.NewInlineKeyboardButtonData("➡️ ", "view-right-task"),
	),
)
