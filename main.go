package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rahmaniali-ir/todo-server/api"
	"github.com/rahmaniali-ir/todo-server/todo"
)

func handlePreFlight(w http.ResponseWriter, r *http.Request) bool {
	if(r.Method != "OPTIONS") {
		return false
	}

	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.WriteHeader(http.StatusOK)
	return true
}

func main() {
	todos := todo.NewCollection("./db/todo.db")

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")

		res := api.ApiResponse{
			Success: true,
			Body: todos.ToArray(),
		}
		res.RespondJSON(w, 200)
	})

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		preflight := handlePreFlight(w, r)
		if preflight {
			return
		}
		
		switch r.Method {
		case "POST":
			
			todo := todo.Todo{}
			err := json.NewDecoder(r.Body).Decode(&todo)
			if err == nil {}

			res := api.ApiResponse{
				Success: true,
				Body: todos.AddTodo(todo),
			}
			res.RespondJSON(w, 200)

		case "DELETE":
			uid := r.URL.Query()["uid"][0]

			todos.DeleteTodo(uid)

			res := api.ApiResponse{
				Success: true,
			}
			res.RespondJSON(w, 200)

		case "PUT":
			uid := r.URL.Query()["uid"][0]

			todo := todos.ToggleTodo(uid)

			res := api.ApiResponse{
				Success: true,
				Body: todo,
			}
			res.RespondJSON(w, 200)
		}
	})

	http.HandleFunc("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		preflight := handlePreFlight(w, r)
		if preflight {
			return
		}
		
		switch r.Method {
		case "POST":
			var credentials struct{
				Email string `json:"email"`
				Password string `json:"password"`
			}

			err := json.NewDecoder(r.Body).Decode(&credentials)
			if err == nil {}

			fmt.Println(credentials)
			
			res := api.ApiResponse{
				Success: true,
				Body: "users.ToArray()",
			}
			res.RespondJSON(w, 200)
		}
	})

	http.ListenAndServe(":8081", nil)
}
