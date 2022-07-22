package todo

import (
	"net/http"
)

type IHandler interface {
	ListAll(*http.Request) (interface{}, error)
}
