package todo

type Status int
const (
	Status_Undone Status = iota
	Status_In_Progress
	Status_Done
)

type Todo struct {
	Uid string `json:"id"`
  Title string `json:"title"`
  Body string `json:"body"`
  Status Status `json:"status"`
	User_uid string `json:"user_uid"`
}

type ITodo interface {
	GetAll() ([]Todo, error)
	GetUserTodos(userUid string) ([]Todo, error)
	AddTodo(todo *Todo) error
}
