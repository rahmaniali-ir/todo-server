package routes

import (
	"net/http"

	"github.com/rahmaniali-ir/todo-server/handlers/todo"
	internalHttp "github.com/rahmaniali-ir/todo-server/internal/http"
	"github.com/rahmaniali-ir/todo-server/router"
)

func TodoRoutes(todoHandler todo.IHandler) []router.Route {
	return []router.Route{
		{
			Name: "getTodos",
			Path: "/todos",
			Method: http.MethodGet,
			Handler: internalHttp.Handle(todoHandler.ListAll),
		},
	}
}
