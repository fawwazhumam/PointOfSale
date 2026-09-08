package request

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Photo    string `json:"photo"`
	Phone    string `json:"phone"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
	Phone    string `json:"phone"`
}

type AssignUserToRoleRequest struct {
	UserID uint `json:"user_id" validate:"required"`
	RoleID uint `json:"role_id" validate:"required"`
}

type EditAssignUserToRoleRequest struct {
	UserID uint `json:"user_id" validate:"required"`
	RoleID uint `json:"role_id" validate:"required"`
}

type GetAllUserRolesRequest struct {
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
	Search    string `query:"search"`
	SortBy    string `query:"sort_by"`
	SortOrder string `query:"sort_order"`
}
