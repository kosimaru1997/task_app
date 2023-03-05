package usecase

import (
	"github.com/kosimaru1997/task_app/app/request"
	"github.com/kosimaru1997/task_app/db/entity"
	"github.com/kosimaru1997/task_app/db/repository"
)

// TodoUseCase はTodoアプリのUseCaseを表します。
type TodoUseCase struct {
	todoRepository *repository.TodoRepository
}

// NewTodoUseCase は新しいTodoUseCaseのインスタンスを返します。
func NewTodoUseCase(todoRepository *repository.TodoRepository) *TodoUseCase {
	return &TodoUseCase{todoRepository: todoRepository}
}

// CreateTodo は新しいTodoを作成します。
func (u *TodoUseCase) CreateTodo(req request.CreateTodoRequest) (*entity.Todo, error) {
	todo := entity.NewTodo(req)
	return u.todoRepository.Create(todo)
}

// GetAllTodos は全てのTodoを取得します。
func (u *TodoUseCase) GetAllTodos() ([]*entity.Todo, error) {
	return u.todoRepository.FindAll()
}

// GetTodoByID は指定されたIDのTodoを取得します。
func (u *TodoUseCase) GetTodoByID(id uint64) (*entity.Todo, error) {
	return u.todoRepository.FindByID(id)
}

// UpdateTodo は指定されたIDのTodoを更新します。
func (u *TodoUseCase) UpdateTodo(id uint64, req request.UpdateTodoRequest) (*entity.Todo, error) {
	todo, err := u.todoRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	todo.Update(req)
	return u.todoRepository.Update(todo)
}

// DeleteTodo は指定されたIDのTodoを削除します。
func (u *TodoUseCase) DeleteTodo(id uint64) error {
	return u.todoRepository.Delete(id)
}
