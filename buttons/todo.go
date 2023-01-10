package buttons

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Todo_view(id, action string) tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("✅ Done", "done!"+id+"!"+action),
			tg.NewInlineKeyboardButtonData("🗑 Delete", "delete!"+id+"!"+action),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("🔙 Back", "back!"+"!"+action),
		),
	)
}
func Todo_Done(id, action string) tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("🔄 Restore", "restore!"+id+"!"+action),
			tg.NewInlineKeyboardButtonData("🗑 Delete", "delete!"+id+"!"+action),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("🔙 Back", "back!"+"!"+action),
		),
	)
}
