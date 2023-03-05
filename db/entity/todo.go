package entity

import (
	"time"

	"github.com/kosimaru1997/task_app/app/request"
)

// Todo はTodoアイテムを表します。
type Todo struct {
	ID        uint64 `gorm:"primaryKey"`
	Title     string
	Memo      string
	DueDate   time.Time `json:"dueDate"`
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewTodo は新しいTodoのインスタンスを返します。
func NewTodo(req request.CreateTodoRequest) *Todo {
	return &Todo{
		Title:     req.Title,
		Memo:      req.Memo,
		DueDate:   req.DueDate.Date,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Complete はTodoを完了状態にします。
func (t *Todo) Complete() {
	t.Completed = true
	t.UpdatedAt = time.Now()
}

// Update はTodoを更新します。
func (t *Todo) Update(req request.UpdateTodoRequest) {
	t.Title = req.Title
	t.Memo = req.Memo
	t.DueDate = req.DueDate.Date
    t.Completed = req.Completed
	t.UpdatedAt = time.Now()
}
