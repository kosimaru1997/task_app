package repository

import (
	"github.com/kosimaru1997/task_app/config"
	"github.com/kosimaru1997/task_app/db/entity"
	"gorm.io/gorm"
)

// TodoRepository はTodoリポジトリを表します。
type TodoRepository struct {
	db *gorm.DB
}

// NewTodoRepository は新しいTodoRepositoryのインスタンスを返します。
func NewTodoRepository() *TodoRepository {
	db, _ := config.ConnectDB()
	return &TodoRepository{db: db}
}

// Create は新しいTodoを作成します。
func (r *TodoRepository) Create(todo *entity.Todo) (*entity.Todo, error) {
	if err := r.db.Create(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

// FindAll は全てのTodoを取得します。
func (r *TodoRepository) FindAll() ([]*entity.Todo, error) {
	todos := []*entity.Todo{}
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// FindByID は指定されたIDのTodoを取得します。
func (r *TodoRepository) FindByID(id uint64) (*entity.Todo, error) {
	todo := &entity.Todo{}
	if err := r.db.First(todo, id).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

// Update は指定されたTodoを更新します。
func (r *TodoRepository) Update(todo *entity.Todo) (*entity.Todo, error) {
	if err := r.db.Save(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

// Delete は指定されたIDのTodoを削除します。
func (r *TodoRepository) Delete(id uint64) error {
	todo := &entity.Todo{}
	if err := r.db.First(todo, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(todo).Error; err != nil {
		return err
	}
	return nil
}
