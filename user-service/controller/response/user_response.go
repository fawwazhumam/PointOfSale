package response

import "pos/user-service/pkg/pagination"

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Photo     string    `json:"photo"`
	Phone     string    `json:"phone"`
	RoleName     string  `json:"role_name"`
}

type GetAllUsersResponse struct {
	Users      []UserResponse                `json:"users"`
	Pagination pagination.PaginationResponse `json:"pagination"`
}


