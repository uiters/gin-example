package form

type UserRoleForm struct {
	Username string   `json:"username"`
	Role     string   `json:"role"`
	Access   []string `json:"access"`
}
