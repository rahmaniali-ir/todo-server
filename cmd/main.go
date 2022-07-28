package main

import (
	"github.com/rahmaniali-ir/todo-server/app"
)

func main() {
	appServer, err := app.New()
	if err != nil {
		panic(err)
	}

	err = appServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
