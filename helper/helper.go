package helper

import (
	"fmt"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/realtemirov/projects/tgbot/model"
	"github.com/spf13/cast"
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
	t2 = time.Hour * time.Duration(cast.ToInt(t))
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

func SliceToInlineKeyboard(todos []*model.Todo, action string) tg.InlineKeyboardMarkup {

	buttons := tg.NewInlineKeyboardMarkup()
	count := 0
	for {
		if len(todos)-count >= 8 {
			buttons.InlineKeyboard = append(buttons.InlineKeyboard, tg.NewInlineKeyboardRow(
				tg.NewInlineKeyboardButtonData(cast.ToString(count+1), "id!"+todos[count].ID+"!"+action),
				tg.NewInlineKeyboardButtonData(cast.ToString(count+2), "id!"+todos[count+1].ID+"!"+action),
				tg.NewInlineKeyboardButtonData(cast.ToString(count+3), "id!"+todos[count+2].ID+"!"+action),
				tg.NewInlineKeyboardButtonData(cast.ToString(count+4), "id!"+todos[count+3].ID+"!"+action),
				tg.NewInlineKeyboardButtonData(cast.ToString(count+5), "id!"+todos[count+4].ID+"!"+action),
				tg.NewInlineKeyboardButtonData(cast.ToString(count+6), "id!"+todos[count+5].ID+"!"+action),
				tg.NewInlineKeyboardButtonData(cast.ToString(count+7), "id!"+todos[count+6].ID+"!"+action),
				tg.NewInlineKeyboardButtonData(cast.ToString(count+8), "id!"+todos[count+7].ID+"!"+action),
			))
			count += 8
		} else {
			switch len(todos) - count {
			case 1:
				{
					buttons.InlineKeyboard = append(buttons.InlineKeyboard, tg.NewInlineKeyboardRow(
						tg.NewInlineKeyboardButtonData(cast.ToString(count+1), "id!"+todos[count].ID+"!"+action),
					))
				}
			case 2:
				{
					buttons.InlineKeyboard = append(buttons.InlineKeyboard, tg.NewInlineKeyboardRow(
						tg.NewInlineKeyboardButtonData(cast.ToString(count+1), "id!"+todos[count].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+2), "id!"+todos[count+1].ID+"!"+action),
					))
				}
			case 3:
				{
					buttons.InlineKeyboard = append(buttons.InlineKeyboard, tg.NewInlineKeyboardRow(
						tg.NewInlineKeyboardButtonData(cast.ToString(count+1), "id!"+todos[count].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+2), "id!"+todos[count+1].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+3), "id!"+todos[count+2].ID+"!"+action),
					))
				}
			case 4:
				{
					buttons.InlineKeyboard = append(buttons.InlineKeyboard, tg.NewInlineKeyboardRow(
						tg.NewInlineKeyboardButtonData(cast.ToString(count+1), "id!"+todos[count].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+2), "id!"+todos[count+1].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+3), "id!"+todos[count+2].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+4), "id!"+todos[count+3].ID+"!"+action),
					))
				}
			case 5:
				{
					buttons.InlineKeyboard = append(buttons.InlineKeyboard, tg.NewInlineKeyboardRow(
						tg.NewInlineKeyboardButtonData(cast.ToString(count+1), "id!"+todos[count].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+2), "id!"+todos[count+1].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+3), "id!"+todos[count+2].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+4), "id!"+todos[count+3].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+5), "id!"+todos[count+4].ID+"!"+action),
					))
				}
			case 6:
				{
					buttons.InlineKeyboard = append(buttons.InlineKeyboard, tg.NewInlineKeyboardRow(
						tg.NewInlineKeyboardButtonData(cast.ToString(count+1), "id!"+todos[count].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+2), "id!"+todos[count+1].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+3), "id!"+todos[count+2].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+4), "id!"+todos[count+3].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+5), "id!"+todos[count+4].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+6), "id!"+todos[count+5].ID+"!"+action),
					))
				}
			case 7:
				{
					buttons.InlineKeyboard = append(buttons.InlineKeyboard, tg.NewInlineKeyboardRow(
						tg.NewInlineKeyboardButtonData(cast.ToString(count+1), "id!"+todos[count].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+2), "id!"+todos[count+1].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+3), "id!"+todos[count+2].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+4), "id!"+todos[count+3].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+5), "id!"+todos[count+4].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+6), "id!"+todos[count+5].ID+"!"+action),
						tg.NewInlineKeyboardButtonData(cast.ToString(count+7), "id!"+todos[count+6].ID+"!"+action),
					))
				}
			}
			break
		}
	}

	return buttons
}
