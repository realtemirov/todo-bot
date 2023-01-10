package model

import (
	"fmt"
	"strings"
	"time"
)

type Base struct {
	ID        string     `json:"id" db:"id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type User struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Action    string    `json:"action" db:"action" default:"empty"`
}

type Todo struct {
	Base
	User_ID      int64      `json:"user_id" db:"user_id"`
	Text         string     `json:"text" db:"text"`
	Photo_URL    string     `json:"photo_url" db:"photo_url"`
	File_URL     string     `json:"file_url" db:"file_url"`
	Deadline     *time.Time `json:"deadline" db:"deadline"`
	Is_Set       bool       `json:"is_set" db:"is_set" default:"false"`
	Is_Done      bool       `json:"is_done" db:"is_done" default:"false"`
	Notification *time.Time `json:"notification" db:"notification"`
}

type Notification struct {
	Time *time.Time `json:"notification" db:"notification"`
	ID   string     `json:"id" db:"id"`
}

func (t *Todo) ToString() string {
	var time time.Time
	txt := fmt.Sprintf("<b>Text:</b>   <i>%s</i>", strings.Join(strings.Split(t.Text, "!"), "\n"))

	if t.Deadline != nil {
		if t.Deadline.Year() != time.Year() {
			txt += "<b>📅 Deadline:</b> <code>" + t.Deadline.Format("15:04 02.01.2006") + "</code>\n"
			t.Notification = nil
		}

	}

	if t.Is_Set {
		if t.Is_Done {
			txt += "<b>✅ Done</b>\n"
		} else {
			txt += "<b>❌ Not done</b>\n"
		}
	}

	if t.Notification != nil {
		if *t.Notification != time {
			txt += "<b>🔔 Notification:</b> <code>" + t.Notification.Format("15:04") + "</code>\n"
		}
	}
	return txt + "\n"
}
