package response

import (
	"time"

	"github.com/kosimaru1997/task_app/db/entity"
)

// TodoResponse はTodoのレスポンスを表します。
type TodoResponse struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Memo      string    `json:"memo"`
	DueDate   string    `json:"due_date"`
	Completed bool      `json:"competed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewTodoResponse はTodoドメインからレスポンスを生成します。
func NewTodoResponse(todo *entity.Todo) *TodoResponse {
	return &TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Memo:      todo.Memo,
		DueDate:   todo.DueDate.Format("2006-01-02"),
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

// TodoListResponse はTodoのリストレスポンスを表します。
type TodoListResponse struct {
	Todos []*TodoResponse `json:"todos"`
}

// NewTodoListResponse はTodoのリストからレスポンスを生成します。
func NewTodoListResponse(todos []*entity.Todo) *TodoListResponse {
	var todoResponses []*TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, NewTodoResponse(todo))
	}
	return &TodoListResponse{Todos: todoResponses}
}
