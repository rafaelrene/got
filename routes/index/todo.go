package index

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	CreatedAt time.Time `field:"created_at"`
	UpdatedAt time.Time `field:"updated_at"`
	Title     string    `field:"title"`
	Id        uuid.UUID `field:"id"`
	IsDone    bool      `field:"is_done"`
}

func NewTodo(title string) Todo {
	return Todo{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     title,
		Id:        uuid.New(),
		IsDone:    false,
	}
}
