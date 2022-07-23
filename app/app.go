package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	todoHandler "github.com/rahmaniali-ir/todo-server/handlers/todo"
	userHandler "github.com/rahmaniali-ir/todo-server/handlers/user"
	todoModel "github.com/rahmaniali-ir/todo-server/models/todo"
	userModel "github.com/rahmaniali-ir/todo-server/models/user"
	"github.com/rahmaniali-ir/todo-server/pkg/session"
	"github.com/rahmaniali-ir/todo-server/router"
	"github.com/rahmaniali-ir/todo-server/routes"
	todoService "github.com/rahmaniali-ir/todo-server/services/todo"
	userService "github.com/rahmaniali-ir/todo-server/services/user"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	dbPath = "./db"
	port = 8081
)

type app struct {
	router *mux.Router
}

func New() (*http.Server, error) {
	allRoutes := []router.Route{}

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic("Could not open database!")
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
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: a.router,
	}
	
	fmt.Printf("Listening on port: %v\n", port)
	return server
}
