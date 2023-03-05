package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kosimaru1997/task_app/app/handler"
	"github.com/kosimaru1997/task_app/app/usecase"
	"github.com/kosimaru1997/task_app/db/repository"
)

func setTodoRoutes(api *gin.RouterGroup) {
	// Todoリポジトリを作成する。
	todoRepo := repository.NewTodoRepository()
	// TodoUseCaseを作成する。
	todoUseCase := usecase.NewTodoUseCase(todoRepo)
	// TodoHandlerを作成する。
	todoHandler := handler.NewTodoHandler(todoUseCase)

	todo := api.Group("/todo")
	{
		todo.GET("", todoHandler.GetAllTodos)
		todo.POST("", todoHandler.CreateTodo)
		todo.GET("/:id", todoHandler.GetTodoByID)
		todo.PUT("/:id", todoHandler.UpdateTodo)
		todo.DELETE("/:id", todoHandler.DeleteTodo)
	}
}

func setRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		setTodoRoutes(api)
	}
}
