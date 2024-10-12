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

func (h IndexHandler) HandleToggleTodoState(c echo.Context) error {
	id := c.FormValue("id")

	if len(id) == 0 {
		fmt.Fprintf(os.Stderr, "Provided `id` is empty string: %s\n", id)
		os.Exit(1)
	}

	todo := h.fetchTodo(id)
	todo.IsDone = !todo.IsDone

	_, err := h.Db.Query("UPDATE todos SET is_done = ? WHERE id = ?", todo.IsDone, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating a TODO in database: %v\n", err)
		os.Exit(1)
	}

	return routes.Render(c, RenderTodo(todo))
}

func (h IndexHandler) fetchAllTodos() []Todo {
	rows, err := h.Db.Query("SELECT * FROM todos;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	return ParseTodos(rows)
}

func (h IndexHandler) fetchTodo(id string) Todo {
	rows, err := h.Db.Query("SELECT * FROM todos WHERE id = ? LIMIT 1;", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	todos := ParseTodos(rows)

	if len(todos) == 0 {
		fmt.Fprintf(os.Stderr, "Failed to find todo with id '%s': %v\n", id, err)
		os.Exit(1)
	}

	return todos[0]
}
