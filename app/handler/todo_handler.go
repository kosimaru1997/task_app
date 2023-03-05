// app/handler/todo_handler.go
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kosimaru1997/task_app/app/request"
	"github.com/kosimaru1997/task_app/app/response"
	"github.com/kosimaru1997/task_app/app/usecase"
)

// TodoHandler はTodoアプリのHandlerを表します。
type TodoHandler struct {
	todoUseCase *usecase.TodoUseCase
}

// NewTodoHandler は新しいTodoHandlerのインスタンスを返します。
func NewTodoHandler(todoUseCase *usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{todoUseCase: todoUseCase}
}

// CreateTodo は新しいTodoを作成します。
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req request.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.todoUseCase.CreateTodo(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var res = response.NewTodoResponse(todo)
	c.JSON(http.StatusCreated, res)
}

// GetAllTodos は全てのTodoを取得します。
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.todoUseCase.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewTodoListResponse(todos))
}

// GetTodoByID は指定されたIDのTodoを取得します。
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	todo, err := h.todoUseCase.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if todo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	c.JSON(http.StatusOK, response.NewTodoResponse(todo))
}

// UpdateTodo は指定されたIDのTodoを更新します。
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req request.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.todoUseCase.UpdateTodo(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewTodoResponse(todo))
}

// DeleteTodo は指定されたIDのTodoを削除します。
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.todoUseCase.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
