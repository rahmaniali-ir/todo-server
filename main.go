package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

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

	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.WriteHeader(http.StatusOK)
	return true
}

func getAuthHeaderToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	headerParts := strings.Fields(authHeader)
	return headerParts[1]
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

	getUserByUid := func (uid string) (user.User, error) {
		dbUserBytes, err := usersDB.Get([]byte(uid), nil)
		if err != nil {
			return user.User{}, errors.New("Invalid user uid!")
		}

		dbUser := user.User{}
		reader := bytes.NewReader(dbUserBytes)
		err = gob.NewDecoder(reader).Decode(&dbUser)
		if err == nil {}

		return dbUser, nil
	}

	getUserUidByToken := func (token string) (string, error) {
		uidBytes, err := tokenDB.Get([]byte(token), nil)
		if err != nil {
			return "", errors.New("Invalid user uid!")
		}

		return string(uidBytes), nil
	}

	// getUserByToken := func (token string) (user.User, error) {
	// 	uid, err := getUserUidByToken(token)
	// 	if err != nil {
	// 		return user.User{}, errors.New("Invalid user token!")
	// 	}

	// 	return getUserByUid(uid)
	// }

	getUserByHeaderToken := func (r *http.Request) (user.User, error) {
		token := getAuthHeaderToken(r)

		uid, err := getUserUidByToken(token)
		if err != nil {
			return user.User{}, errors.New("Invalid user token!")
		}

		return getUserByUid(uid)
	}

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")

		preflight := handlePreFlight(w, r)
		if preflight {
			return
		}

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

		preflight := handlePreFlight(w, r)
		if preflight {
			return
		}
		
		dbUser, err := getUserByHeaderToken(r)
		if err != nil {
			res := api.ApiResponse{
				Success: false,
				Body: nil,
			}
			res.RespondJSON(w, 404)
			return
		}

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

	http.HandleFunc("/sign-out", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		preflight := handlePreFlight(w, r)
		if preflight {
			return
		}
		
		switch r.Method {
		case "DELETE":
			token := getAuthHeaderToken(r)

			_, err := getUserByHeaderToken(r)
			if err != nil {}

			err = tokenDB.Delete([]byte(token), nil)
			if err != nil {
				res := api.ApiResponse{
					Success: false,
					Body: nil,
					Message: "Invalid token!",
				}
				res.RespondJSON(w, 400)
				return
			}
			
			res := api.ApiResponse{
				Success: true,
				Body: nil,
			}
			res.RespondJSON(w, 200)
		}
	})

	http.ListenAndServe(":8081", nil)
}
