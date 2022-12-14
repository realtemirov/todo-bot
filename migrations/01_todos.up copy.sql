CREATE TABLE IF NOT EXISTS "notification" (
    "id" text PRIMARY KEY,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp,
    "deleted_at" timestamp,
    "user_id" bigint NOT NULL,
    "todo_id" text NOT NULL,
    "notif_date" timestamp,
    CONSTRAINT fk_todo
      FOREIGN KEY(todo_id) 
	  REFERENCES todos(id)
);

-- type Notification struct {
-- 	Base
-- 	Todo_ID    string    `json:"todo_id" db:"todo_id"`
-- 	User_ID    int64     `json:"user_id" db:"user_id"`
-- 	Notif_date time.Time `json:"notif_date" db:"notif_date"`
-- }
