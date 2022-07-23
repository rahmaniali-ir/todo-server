package http

import "net/http"

type handlerFunc func(handler *GenericRequest) (interface{}, error)

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

		res := GenericResponse{}

		gr, err := NewGenericResponseFromHTTPRequest(r)
		if err != nil {
			res.RespondJSON(w, 400)
			return
		}

		body, err := handler(gr)
		if err != nil {
			return
		}

		res = GenericResponse{
			Success: true,
			Body: body,
		}
		res.RespondJSON(w, http.StatusOK)
	}
}
