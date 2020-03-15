package form

type ToDoForm struct {
	Name string `json:"name" binding:"required"`
}
