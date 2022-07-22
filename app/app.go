package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/rahmaniali-ir/todo-server/handlers/todo"
	model "github.com/rahmaniali-ir/todo-server/models/todo"
	"github.com/rahmaniali-ir/todo-server/router"
	"github.com/rahmaniali-ir/todo-server/routes"
	service "github.com/rahmaniali-ir/todo-server/services/todo"
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

	// todo routes
	todoModel, err := model.NewModel(db)
	todoService := service.NewService(todoModel)
	allRoutes = append(allRoutes, routes.TodoRoutes(handler.NewHandler(&todoService))...)

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
