package index

import (
	"database/sql"
	"fmt"
	"os"
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

func ParseTodos(rows *sql.Rows) []Todo {
	var todos []Todo

	for rows.Next() {
		var todo Todo
		var createdAt string
		var updatedAt string

		if err := rows.Scan(&todo.Id, &todo.Title, &todo.IsDone, &createdAt, &updatedAt); err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning row: %v\n", err)
			os.Exit(1)
		}

		parsedCreatedAt, err := time.Parse(time.RFC3339, createdAt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't parse created_at (%s): %v\n", createdAt, err)
			os.Exit(1)
		}
		todo.CreatedAt = parsedCreatedAt

		parsedUpdatedAt, err := time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't parse updated_at (%s): %v\n", updatedAt, err)
			os.Exit(1)
		}
		todo.UpdatedAt = parsedUpdatedAt

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error during rows iteration: %v\n", err)
		os.Exit(1)
	}

	return todos
}
