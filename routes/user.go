package routes

import (
	"net/http"

	"github.com/rahmaniali-ir/todo-server/handlers/user"
	internalHttp "github.com/rahmaniali-ir/todo-server/internal/http"
	"github.com/rahmaniali-ir/todo-server/router"
)

func UserRoutes(userHandler user.IHandler) []router.Route {
	return []router.Route{
		{
			Name: "getProfile",
			Path: "/profile",
			Method: http.MethodGet,
			Handler: internalHttp.Handle(userHandler.GetProfile),
		},
		{
			Name: "signUp",
			Path: "/sign-up",
			Method: http.MethodPost,
			Handler: internalHttp.Handle(userHandler.SignUp),
		},
	}
}
