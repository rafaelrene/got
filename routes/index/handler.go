package index

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rafaelrene/got/routes"
)

type IndexHandler struct {
	Db *sql.DB
}

func (h IndexHandler) HandleIndexShow(c echo.Context) error {
	todos := h.fetchAllTodos()

	return routes.Render(c, Show(todos))
}

func (h IndexHandler) HandleAddTodo(c echo.Context) error {
	title := c.FormValue("title")
	todo := NewTodo(title)

	_, err := h.Db.Query("INSERT INTO todos(id, title, is_done, created_at, updated_at) VALUES(?, ?, ?, ?, ?)", todo.Id, todo.Title, todo.IsDone, todo.CreatedAt, todo.UpdatedAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing a TODO into databse: %v\n", err)
		os.Exit(1)
	}

	return routes.Render(c, AddTodo(todo))
}

func (h IndexHandler) fetchAllTodos() []Todo {
	rows, err := h.Db.Query("SELECT * FROM todos;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo

		if err := rows.Scan(&todo.Id, &todo.Title, &todo.IsDone, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning row: %v\n", err)
			os.Exit(1)
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error during rows iteration: %v\n", err)
		os.Exit(1)
	}

	return todos
}
