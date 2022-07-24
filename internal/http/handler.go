package http

import (
	"net/http"
)

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
			res = GenericResponse{
				Success: false,
				Body: err.Error(),
			}
			res.RespondJSON(w, http.StatusBadRequest)
			return
		}

		body, err := handler(gr)
		if err != nil {
			res = GenericResponse{
				Success: false,
				Message: err.Error(),
			}
			res.RespondJSON(w, http.StatusBadRequest)

			return
		}

		res = GenericResponse{
			Success: true,
			Body: body,
		}
		res.RespondJSON(w, http.StatusOK)
	}
}
