package request

import (
	"github.com/kosimaru1997/task_app/common"
)

// CreateTodoRequest はTodo作成APIのリクエストを表します。
type CreateTodoRequest struct {
	Title   string          `json:"title" binding:"required"`
	Memo    string          `json:"memo"`
	DueDate common.OnlyDate `json:"due_date"`
}

// UpdateTodoRequest はTodo更新APIのリクエストを表します。
type UpdateTodoRequest struct {
	Title     string          `json:"title" binding:"required"`
	Memo      string          `json:"memo"`
	Completed bool            `json:"completed"`
	DueDate   common.OnlyDate `json:"due_date"`
}
