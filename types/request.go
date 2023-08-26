package types

type CreateTodoRequest struct {
	Title       string
	Description string
	Deadline    string
	Completed   bool
}
