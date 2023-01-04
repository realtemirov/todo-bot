package helper

import (
	"time"

	"github.com/realtemirov/projects/tgbot/model"
)

func TodoToString(t *model.Todo) string {
	text := ""
	time := time.Time{}
	if t.Base.ID != "" {
		text += "*ID:* __" + t.Base.ID + "__ \n"
	}
	if t.Base.CreatedAt != time {
		text += "*Created Time :* __" + t.Base.CreatedAt.String() + "__ \n"
	}
	if t.Title != "" {
		text += "*Title:* __" + t.Title + "__ \n"
	}
	if t.Description != "" {
		text += "*Description:* __" + t.Description + "__ \n"
	}
	if t.Photo_URL != "" {
		text += "*Photo_URL:* __" + t.Photo_URL + "__ \n"
	}
	if t.File_URL != "" {
		text += "*File_URL:* __" + t.File_URL + "__ \n"
	}
	if t.IsDone {
		text += "*Done:* __✅__  \n"
	}

	return text

}
