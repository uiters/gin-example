package form

type RoleForm struct {
	Name string `json:"name" binding:"required"`
}
