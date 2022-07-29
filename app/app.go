package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	todoHandler "github.com/rahmaniali-ir/todo-server/internal/handlers/todo"
	userHandler "github.com/rahmaniali-ir/todo-server/internal/handlers/user"
	todoModel "github.com/rahmaniali-ir/todo-server/internal/models/todo"
	userModel "github.com/rahmaniali-ir/todo-server/internal/models/user"
	"github.com/rahmaniali-ir/todo-server/internal/router"
	"github.com/rahmaniali-ir/todo-server/internal/routes"
	todoService "github.com/rahmaniali-ir/todo-server/internal/services/todo"
	userService "github.com/rahmaniali-ir/todo-server/internal/services/user"
	"github.com/rahmaniali-ir/todo-server/pkg/session"
	"github.com/syndtr/goleveldb/leveldb"
)

var EnvMap map[string]string
var defaultEnv = map[string]string{
	"SERVER_PORT": "8081",
	"DB_PATH": "./db",
}

type app struct {
	router *mux.Router
}

func New() (*http.Server, error) {
	var err error

	EnvMap, err = godotenv.Read(".env")
	if err != nil {
		EnvMap = defaultEnv
	}

	allRoutes := []router.Route{}

	dbPath := EnvMap["DB_PATH"]
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(err)
	}

	// user routes
	um := userModel.NewModel(db)
	us := userService.NewService(um)
	allRoutes = append(allRoutes, routes.UserRoutes(userHandler.NewHandler(us))...)

	// todo routes
	tm, err := todoModel.NewModel(db)
	ts := todoService.NewService(tm)
	allRoutes = append(allRoutes, routes.TodoRoutes(todoHandler.NewHandler(&ts))...)

	// session manager
	session.Init(db)

	newApp := app{}
	err = newApp.createResources(allRoutes...)

	if err != nil {
		return nil, err
	}

	return newApp.server(), nil
}

func (a *app) createResources(rs ...router.Route) error {
	a.router = mux.NewRouter().StrictSlash(true)

	for _, r := range rs {
		err := a.router.Name(r.Name).Path(r.Path).Methods(r.Method, http.MethodOptions).HandlerFunc(r.Handler).GetError()

		if err != nil {
			return err
		}
	}

	return nil
}

func (a *app) server() *http.Server {
	port := EnvMap["SERVER_PORT"]

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: a.router,
	}
	
	fmt.Printf("Listening on port: %v\n", port)
	return server
}
