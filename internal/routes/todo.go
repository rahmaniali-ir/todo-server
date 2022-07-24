package routes

import (
	"net/http"

	"github.com/rahmaniali-ir/todo-server/internal/handlers/todo"
	internalHttp "github.com/rahmaniali-ir/todo-server/internal/http"
	"github.com/rahmaniali-ir/todo-server/internal/router"
)

func TodoRoutes(todoHandler todo.IHandler) []router.Route {
	return []router.Route{
		{
			Name: "getTodos",
			Path: "/todos",
			Method: http.MethodGet,
			Handler: internalHttp.Handle(todoHandler.GetAll),
		},
		{
			Name: "addTodo",
			Path: "/todo",
			Method: http.MethodPost,
			Handler: internalHttp.Handle(todoHandler.Add),
		},
		{
			Name: "deleteTodo",
			Path: "/todo",
			Method: http.MethodDelete,
			Handler: internalHttp.Handle(todoHandler.Delete),
		},
	}
}
