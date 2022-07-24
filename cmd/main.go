package main

import (
	"github.com/rahmaniali-ir/todo-server/app"
	// todoController "github.com/rahmaniali-ir/todo-server/controllers/todo"
	// "github.com/rahmaniali-ir/todo-server/todo"
	// "github.com/rahmaniali-ir/todo-server/user"
)

func main() {
	// todos := todo.NewCollection("./db/todo.db")
	// usersDB := user.NewCollection()
	// defer usersDB.Close()

	// todoHandler := todoController.New(&todos, &usersDB)

	appServer, err := app.New()
	if err != nil {
		panic(err)
	}

	err = appServer.ListenAndServe()
	if err != nil {
		panic(err)
	}

	// getUserByToken := func (token string) (user.User, error) {
	// 	uid, err := getUserUidByToken(token)
	// 	if err != nil {
	// 		return user.User{}, errors.New("Invalid user token!")
	// 	}

	// 	return getUserByUid(uid)
	// }
	// http.HandleFunc("/todo", todoHandler.HandleTodo)
	// http.HandleFunc("/todos", todoHandler.HandleTodos)
	
	// http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("Access-Control-Allow-Origin", "*")
	// 	w.Header().Add("Content-Type", "application/json")

	// 	preflight := utils.HandlePreFlight(w, r)
	// 	if preflight {
	// 		return
	// 	}
		
	// 	dbUser, err := usersDB.GetUserByHeaderToken(r)
	// 	if err != nil {
	// 		res := api.ApiResponse{
	// 			Success: false,
	// 			Body: nil,
	// 		}
	// 		res.RespondJSON(w, 404)
	// 		return
	// 	}

	// 	res := api.ApiResponse{
	// 		Success: true,
	// 		Body: dbUser,
	// 	}
	// 	res.RespondJSON(w, 200)
	// })

	// http.HandleFunc("/sign-in", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("Access-Control-Allow-Origin", "*")

	// 	preflight := utils.HandlePreFlight(w, r)
	// 	if preflight {
	// 		return
	// 	}
		
	// 	switch r.Method {
	// 	case "POST":
	// 		var credentials struct{
	// 			Email string `json:"email"`
	// 			Password string `json:"password"`
	// 		}

	// 		err := json.NewDecoder(r.Body).Decode(&credentials)
	// 		if err == nil {}

	// 		foundUser, err := usersDB.SearchSingleUser(func (dbUser user.User) bool {
	// 			return dbUser.Email == credentials.Email && dbUser.Password == credentials.Password
	// 		})

	// 		fmt.Println("Found", foundUser)

	// 		// invalid user
	// 		if err != nil {
	// 			res := api.ApiResponse{
	// 				Success: false,
	// 				Body: nil,
	// 			}
	// 			res.RespondJSON(w, 404)
	// 			return
	// 		}

	// 		token, err := usersDB.SignUserIn(foundUser.Uid)
	// 		if err != nil {}

	// 		var response struct{
	// 			Token string `json:"token"`
	// 			User user.User `json:"user"`
	// 		}
	// 		response.Token = token
	// 		response.User = foundUser
			
	// 		res := api.ApiResponse{
	// 			Success: true,
	// 			Body: response,
	// 		}
	// 		res.RespondJSON(w, 200)
	// 	}
	// })

	// http.HandleFunc("/sign-up", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("Access-Control-Allow-Origin", "*")

	// 	preflight := utils.HandlePreFlight(w, r)
	// 	if preflight {
	// 		return
	// 	}
		
	// 	switch r.Method {
	// 	case "POST":
	// 		var credentials struct{
	// 			Name string `json:"name"`
	// 			Email string `json:"email"`
	// 			Password string `json:"password"`
	// 		}

	// 		err := json.NewDecoder(r.Body).Decode(&credentials)
	// 		if err == nil {}

	// 		fmt.Println("Sign-up", credentials)

	// 		// create user
	// 		newUserData := user.User{
	// 			Name: credentials.Name,
	// 			Email: credentials.Email,
	// 			Password: credentials.Password,
	// 		}

	// 		newUser, err := usersDB.AddUser(newUserData)
	// 		if err != nil {
	// 			res := api.ApiResponse{
	// 				Success: false,
	// 				Body: nil,
	// 				Message: err.Error(),
	// 			}
	// 			res.RespondJSON(w, 400)
	// 			return
	// 		}

	// 		token, err := usersDB.SignUserIn(newUser.Uid)
	// 		if err != nil {
	// 			res := api.ApiResponse{
	// 				Success: false,
	// 				Body: nil,
	// 				Message: err.Error(),
	// 			}
	// 			res.RespondJSON(w, 400)
	// 			return
	// 		}

	// 		// response
	// 		var response struct{
	// 			Token string `json:"token"`
	// 			User user.User `json:"user"`
	// 		}
	// 		response.Token = token
	// 		response.User = newUser
			
	// 		res := api.ApiResponse{
	// 			Success: true,
	// 			Body: response,
	// 		}
	// 		res.RespondJSON(w, 200)
	// 	}
	// })

	// http.HandleFunc("/sign-out", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("Access-Control-Allow-Origin", "*")

	// 	preflight := utils.HandlePreFlight(w, r)
	// 	if preflight {
	// 		return
	// 	}
		
	// 	switch r.Method {
	// 	case "DELETE":
	// 		token := utils.GetAuthHeaderToken(r)

	// 		_, err := usersDB.GetUserByHeaderToken(r)
	// 		if err != nil {}

	// 		err = usersDB.SignUserOut(token)
	// 		if err != nil {
	// 			res := api.ApiResponse{
	// 				Success: false,
	// 				Body: nil,
	// 				Message: "Invalid token!",
	// 			}
	// 			res.RespondJSON(w, 400)
	// 			return
	// 		}
			
	// 		res := api.ApiResponse{
	// 			Success: true,
	// 			Body: nil,
	// 		}
	// 		res.RespondJSON(w, 200)
	// 	}
	// })

	// http.ListenAndServe(":8081", nil)
}
