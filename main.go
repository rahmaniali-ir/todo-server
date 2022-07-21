package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/rahmaniali-ir/todo-server/api"
	"github.com/rahmaniali-ir/todo-server/todo"
	"github.com/rahmaniali-ir/todo-server/user"
	"github.com/syndtr/goleveldb/leveldb"
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

	usersDB, err := leveldb.OpenFile("./db/user", nil)
	if err != nil {
		panic("Could not open database!")
	}
	defer usersDB.Close()

	tokenDB, err := leveldb.OpenFile("./db/auth-token", nil)
	if err != nil {
		panic("Could not open database!")
	}
	defer tokenDB.Close()

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

	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		
		token := r.Header.Get("Authorization")

		dbUserBytes, err := usersDB.Get([]byte(token), nil)
		if err != nil {
			res := api.ApiResponse{
				Success: false,
				Body: nil,
			}
			res.RespondJSON(w, 404)
			return
		}

		var dbUser user.User
		reader := bytes.NewReader(dbUserBytes)
		err = gob.NewDecoder(reader).Decode(&dbUser)
		if err != nil {}

		res := api.ApiResponse{
			Success: true,
			Body: dbUser,
		}
		res.RespondJSON(w, 200)
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

			iter := usersDB.NewIterator(nil, nil)
			for iter.Next() {
				fmt.Println(string(iter.Key()))

				var dbUser user.User
				reader := bytes.NewReader(iter.Value())
				err = gob.NewDecoder(reader).Decode(&dbUser)
				if err != nil {}

				// valid user
				if dbUser.Email == credentials.Email && dbUser.Password == credentials.Password {
					// generate token
					token := uuid.NewString()
					err = tokenDB.Put([]byte(token), []byte(dbUser.Uid), nil)
					if err != nil {}

					// response
					var response struct{
						Token string `json:"token"`
						User user.User `json:"user"`
					}
					response.Token = token
					response.User = dbUser
					
					res := api.ApiResponse{
						Success: true,
						Body: response,
					}
					res.RespondJSON(w, 200)
					return
				}
			}
			
			// invalid user
			res := api.ApiResponse{
				Success: false,
				Body: nil,
			}
			res.RespondJSON(w, 200)
		}
	})

	http.HandleFunc("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		preflight := handlePreFlight(w, r)
		if preflight {
			return
		}
		
		switch r.Method {
		case "POST":
			var credentials struct{
				Name string `json:"name"`
				Email string `json:"email"`
				Password string `json:"password"`
			}

			err := json.NewDecoder(r.Body).Decode(&credentials)
			if err == nil {}

			fmt.Println("Sign-up", credentials)

			// create user
			uid := uuid.NewString()
			newUser := user.User{
				Uid: uid,
				Name: credentials.Name,
				Email: credentials.Email,
				Password: credentials.Password,
			}

			var userBytes bytes.Buffer
			err = gob.NewEncoder(&userBytes).Encode(newUser)
			if err != nil {}
			
			err = usersDB.Put([]byte(uid), userBytes.Bytes(), nil)
			if err != nil {}

			// generate token
			token := uuid.NewString()
			err = tokenDB.Put([]byte(token), []byte(uid), nil)
			if err != nil {}

			// response
			var response struct{
				Token string `json:"token"`
				User user.User `json:"user"`
			}
			response.Token = token
			response.User = newUser
			
			res := api.ApiResponse{
				Success: true,
				Body: response,
			}
			res.RespondJSON(w, 200)
		}
	})

	http.ListenAndServe(":8081", nil)
}
