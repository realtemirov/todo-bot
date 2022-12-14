package model

import "time"

type Base struct {
	ID        string     `json:"id" db:"id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type User struct {
	ID        int64      `json:"id" db:"id"`
	Username  string     `json:"username" db:"username"`
	FirsName  string     `json:"first_name" db:"first_name"`
	LastName  string     `json:"last_name" db:"last_name"`
	IsAdmin   bool       `json:"is_admin" db:"is_admin"`
	Photo_URL string     `json:"photo_url" db:"photo_url"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type Todo struct {
	Base
	User_ID     int64      `json:"user_id" db:"user_id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Photo_URL   string     `json:"photo_url" db:"photo_url"`
	File_URL    string     `json:"file_url" db:"file_url"`
	Deadline    *time.Time `json:"deadline" db:"deadline"`
	IsDone      bool       `json:"is_done" db:"is_done"`
}

type Notification struct {
	Base
	Todo_ID    string    `json:"todo_id" db:"todo_id"`
	User_ID    int64     `json:"user_id" db:"user_id"`
	Notif_date time.Time `json:"notif_date" db:"notif_date"`
}
