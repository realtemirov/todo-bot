package buttons

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/helper/const/query"
	"github.com/realtemirov/projects/tgbot/helper/const/word"
)

var Todo = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.TODO_ADD, query.TODO_ADD),
		//list of tasks
		tg.NewInlineKeyboardButtonData(word.TODO_TODOS, query.TODO_TODOS),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.TODO_HISTORY, query.TODO_HISTORY),
	),
)

var New_todo = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.TODO_NEW_TITLE, query.TODO_NEW_TITLE),
		tg.NewInlineKeyboardButtonData(word.TODO_NEW_DESCRIPTION, query.TODO_NEW_DESCRIPTION),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.TODO_NEW_PICTURE, query.TODO_NEW_PICTURE),
		tg.NewInlineKeyboardButtonData(word.TODO_NEW_FILE, query.TODO_NEW_FILE),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.TODO_NEW_DEADLINE, query.TODO_NEW_DEADLINE),
		tg.NewInlineKeyboardButtonData(word.TODO_NEW_NOTIFICATON, query.TODO_NEW_NOTIFICATON),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(word.TODO_NEW_SAVE, query.TODO_NEW_SAVE),
		tg.NewInlineKeyboardButtonData(word.TODO_NEW_CANCEL, query.TODO_NEW_CANCEL),
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
