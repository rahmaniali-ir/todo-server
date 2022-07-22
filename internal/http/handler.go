package http

import "net/http"

type handlerFunc func(handler *http.Request) (interface{}, error)

func Handle(handler handlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")

		// preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		body, err := handler(r)
		if err != nil {
			return
		}

		res := ApiResponse{
			Success: true,
			Body: body,
		}
		res.RespondJSON(w, http.StatusOK)
	}
}
