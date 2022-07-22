package todoController

import (
	"encoding/json"
	"net/http"

	"github.com/rahmaniali-ir/todo-server/api"
	"github.com/rahmaniali-ir/todo-server/todo"
	"github.com/rahmaniali-ir/todo-server/user"
	"github.com/rahmaniali-ir/todo-server/utils"
)

type TodoHandler struct {
	todos *todo.Collection
	users *user.Collection
}

func New(todos *todo.Collection, users *user.Collection) TodoHandler {
	return TodoHandler{
		todos: todos,
		users: users,
	}
}

func (handler *TodoHandler) HandleTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	preflight := utils.HandlePreFlight(w, r)
	if preflight {
		return
	}

	dbUser, err := handler.users.GetUserByHeaderToken(r)
	if err != nil {
		res := api.ApiResponse{
			Success: false,
			Message: "Invalid user!",
		}
		res.RespondJSON(w, 404)
		return
	}

	filteredTodos := handler.todos.Filter(func(t todo.Todo) bool {
		return t.User_uid == dbUser.Uid
	})

	res := api.ApiResponse{
		Success: true,
		Body: filteredTodos,
	}
	res.RespondJSON(w, 200)
}

func (handler *TodoHandler) HandleTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	preflight := utils.HandlePreFlight(w, r)
	if preflight {
		return
	}
	
	switch r.Method {
	case "POST":
		dbUser, err := handler.users.GetUserByHeaderToken(r)
		if err != nil {
			res := api.ApiResponse{
				Success: false,
				Message: "Invalid user!",
			}
			res.RespondJSON(w, 404)
		}
		
		todo := todo.Todo{}
		err = json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			res := api.ApiResponse{
				Success: false,
				Message: err.Error(),
			}
			res.RespondJSON(w, 400)
			return
		}

		todo.User_uid = dbUser.Uid
		addedTodo := handler.todos.AddTodo(todo)

		res := api.ApiResponse{
			Success: true,
			Body: addedTodo,
		}
		res.RespondJSON(w, 200)

	case "DELETE":
		uid := r.URL.Query()["uid"][0]

		handler.todos.DeleteTodo(uid)

		res := api.ApiResponse{
			Success: true,
		}
		res.RespondJSON(w, 200)

	case "PUT":
		uid := r.URL.Query()["uid"][0]

		todo := handler.todos.ToggleTodo(uid)

		res := api.ApiResponse{
			Success: true,
			Body: todo,
		}
		res.RespondJSON(w, 200)
	}
}

