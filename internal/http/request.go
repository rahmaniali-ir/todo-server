package http

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rahmaniali-ir/todo-server/pkg/constants"
	"github.com/rahmaniali-ir/todo-server/pkg/session"
)

type GenericRequest struct {
	Session *session.TempUserSession
	Method string
	URL string
	Headers http.Header
	QueryParams url.Values
	Body []byte
	PathParams map[string]string

	r *http.Request
}

func NewGenericResponseFromHTTPRequest(r *http.Request) (*GenericRequest, error) {
	gr := &GenericRequest{}

	var err error
	gr.Headers = r.Header
	gr.URL = r.URL.String()
	gr.PathParams = mux.Vars(r)
	gr.QueryParams = r.URL.Query()
	gr.Method = r.Method

	token := r.Header.Get(constants.HTTPXTokenHeader)
	token = strings.TrimPrefix(token, "Bearer ")
	if token != "" {
		gr.Session, err = session.Default.GetByToken(token)

		if err != nil {
			return nil, err
		}
	}

	reader := bufio.NewReader(r.Body)
	body, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	gr.Body = body

	return gr, nil
}
