package user

import (
	"bytes"
	"encoding/json"

	"github.com/rahmaniali-ir/todo-server/internal/http"
	userModel "github.com/rahmaniali-ir/todo-server/models/user"
	"github.com/rahmaniali-ir/todo-server/pkg/session"
	userService "github.com/rahmaniali-ir/todo-server/services/user"
)

type handler struct {
	service userService.IUser
}

var _ IHandler = &handler{}

func NewHandler(service userService.IUser) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetProfile(req *http.GenericRequest) (interface{}, error) {
	return h.service.Get(req.Session.Uid)
}

func (h *handler) SignUp(req *http.GenericRequest) (interface{}, error) {
	user := userModel.User{}
	reader := bytes.NewReader(req.Body)
	err := json.NewDecoder(reader).Decode(&user)
	if err != nil {
		return nil, err
	}

	addedUser, err := h.service.Add(user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	addedUser.Password = ""

	token, err := session.Default.SetSession(addedUser.Uid)
	if err != nil {
		return nil, err
	}

	var credentials struct{
		Token string `json:"token"`
		User userModel.User `json:"user"`
	}
	credentials.User = *addedUser
	credentials.Token = token

	return credentials, nil
}
