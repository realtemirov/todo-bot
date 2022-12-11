CREATE TABLE IF NOT EXISTS "todos" (
    "id" text PRIMARY KEY,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp,
    "deleted_at" timestamp,
    "user_id" bigint NOT NULL,
    "title" text NOT NULL,
    "description" text,
    "photo_url" text,
    "file_url" text,
    "deadline" timestamp,
    "is_done" boolean NOT NULL DEFAULT false
);

/*type Todo struct {
	Base
	User_ID     string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Photo_URL   string     `json:"photo_url"`
	File_URL    string     `json:"file_url"`
	Deadline    *time.Time `json:"deadline"`
	IsDone      bool       `json:"is_done"`
}*/