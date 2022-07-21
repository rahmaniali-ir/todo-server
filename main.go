package main

import (
	"encoding/json"
	"net/http"

	"github.com/rahmaniali-ir/todo-server/api"
	"github.com/rahmaniali-ir/todo-server/todo"
)

func main() {
	collection := todo.NewCollection("./todo.db")

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")

		res := api.ApiResponse{
			Success: true,
			Body: collection.ToArray(),
		}
		res.RespondJSON(w, 200)
	})

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		
		if r.Method == "OPTIONS" {
			w.Header().Add("Access-Control-Allow-Headers", "content-type")
			w.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
			w.WriteHeader(http.StatusOK)
			return
		} else if r.Method == "POST" {
			
			todo := todo.Todo{}
			err := json.NewDecoder(r.Body).Decode(&todo)
			if err == nil {}

			res := api.ApiResponse{
				Success: true,
				Body: collection.AddTodo(todo),
			}
			res.RespondJSON(w, 200)
		} else if r.Method == "DELETE" {
			uid := r.URL.Query()["uid"][0]

			collection.DeleteTodo(uid)

			res := api.ApiResponse{
				Success: true,
			}
			res.RespondJSON(w, 200)
		} else if r.Method == "PUT" {
			uid := r.URL.Query()["uid"][0]

			todo := collection.ToggleTodo(uid)

			res := api.ApiResponse{
				Success: true,
				Body: todo,
			}
			res.RespondJSON(w, 200)
		}

	})

	http.ListenAndServe(":8081", nil)
}
