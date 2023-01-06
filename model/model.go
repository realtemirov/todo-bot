package model

import (
	"fmt"
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
	Title        string     `json:"title" db:"title"`
	Description  string     `json:"description" db:"description"`
	Photo_URL    string     `json:"photo_url" db:"photo_url"`
	File_URL     string     `json:"file_url" db:"file_url"`
	Deadline     *time.Time `json:"deadline" db:"deadline"`
	Is_Set       bool       `json:"is_set" db:"is_set" default:"false"`
	Notification *time.Time `json:"notification" db:"notification"`
}

type Notification struct {
	Time *time.Time `json:"notification" db:"notification"`
	ID   string     `json:"id" db:"id"`
}

func (t *Todo) ToString() string {
	txt := fmt.Sprintf("<b>Title:<b> <i>%s</i> \n", t.Title)
	if t.Description != "" {
		txt += fmt.Sprintf("<b>Description:<b> <i>%s</i> \n", t.Description)
	}

	return txt
}
