package controller

import (
	"pos/user-service/controller/request"
	"pos/user-service/controller/response"
	"pos/user-service/model"
	"pos/user-service/pkg/conv"
	"pos/user-service/pkg/pagination"
	"pos/user-service/pkg/validator"
	"pos/user-service/usecase"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type UserControllerInterface interface {
	CreateUser(c fiber.Ctx) error
	GetAllUsers(c fiber.Ctx) error
	GetUserByID(c fiber.Ctx) error
	GetUserByEmail(c fiber.Ctx) error
	UpdateUser(c fiber.Ctx) error
	DeleteUser(c fiber.Ctx) error

	GetUserByRoleName(c fiber.Ctx) error

	AssignUserToRole(c fiber.Ctx) error
	EditAssignUserToRole(c fiber.Ctx) error
	GetUserRoleByID(c fiber.Ctx) error
	GetAllUserRoles(c fiber.Ctx) error
}

type userController struct {
	userUsecase usecase.UserUsecaseInterface
}

func NewUserController(userUsecase usecase.UserUsecaseInterface) UserControllerInterface {
	return &userController{userUsecase: userUsecase}
}

// CreateUser implements UserControllerInterface
func (u *userController) CreateUser(c fiber.Ctx) error {
	ctx := c.Context()

	req := request.CreateUserRequest{}

	if err := c.Bind().Body(&req); err != nil {
		log.Errorf("[UserController] CreateUser - 1: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[UserController] CreateUser - 2: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userModel := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Photo:    req.Photo,
		Phone:    req.Phone,
	}

	if err := u.userUsecase.CreateUser(ctx, userModel); err != nil {
		log.Errorf("[UserController] CreateUser - 3: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

// GetAllUsers implements UserControllerInterface
func (u *userController) GetAllUsers(c fiber.Ctx) error {
	ctx := c.Context()

	page := fiber.Query[int](c, "page", 1)
	limit := fiber.Query[int](c, "limit", 10)
	search := fiber.Query[string](c, "search", "")
	sortBy := fiber.Query[string](c, "sort_by", "")
	sortOrder := fiber.Query[string](c, "sort_order", "")

	users, totalRecord, err := u.userUsecase.GetAllUsers(ctx, page, limit, search, sortBy, sortOrder)
	if err != nil {
		log.Errorf("[UserController] GetAllUsers - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resp := []response.UserResponse{}
	for _, user := range users {
		roleName := ""
		if len(user.Roles) > 0 {
			roleName = user.Roles[0].Name
		}

		resp = append(resp, response.UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Photo:    user.Photo,
			Phone:    user.Phone,
			RoleName: roleName,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Users fetched successfully",
		"total":   totalRecord,
		"data":    resp,
	})
}

// GetUserByID implements UserControllerInterface
func (u *userController) GetUserByID(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Params("id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User id is required",
		})
	}

	id := conv.StringToUint(userID)

	user, err := u.userUsecase.GetUserByID(ctx, id)
	if err != nil {
		log.Errorf("[UserController] GetUserByID - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	roleName := ""
	if len(user.Roles) > 0 {
		roleName = user.Roles[0].Name
	}

	resp := response.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Photo:    user.Photo,
		Phone:    user.Phone,
		RoleName: roleName,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User fetched successfully",
		"data":    resp,
	})
}

// GetUserByEmail implements UserControllerInterface
func (u *userController) GetUserByEmail(c fiber.Ctx) error {
	ctx := c.Context()

	email := c.Params("email")
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User email is required",
		})
	}

	user, err := u.userUsecase.GetUserByEmail(ctx, email)
	if err != nil {
		log.Errorf("[UserController] GetUserByEmail - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	roleName := ""
	if len(user.Roles) > 0 {
		roleName = user.Roles[0].Name
	}

	resp := response.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Photo:    user.Photo,
		Phone:    user.Phone,
		RoleName: roleName,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User fetched successfully",
		"data":    resp,
	})
}

// UpdateUser implements UserControllerInterface
func (u *userController) UpdateUser(c fiber.Ctx) error {
	ctx := c.Context()

	req := request.UpdateUserRequest{}
	if err := c.Bind().Body(&req); err != nil {
		log.Errorf("[UserController] UpdateUser - 1: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[UserController] UpdateUser - 2: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reqModel := model.User{
		ID:       conv.StringToUint(c.Params("id")),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Photo:    req.Photo,
		Phone:    req.Phone,
	}

	if req.Password != "" {
		hashedPassword, err := conv.HashPassword(req.Password)
		if err != nil {
			log.Errorf("[UserController] UpdateUser - 3: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		reqModel.Password = hashedPassword
	}

	if err := u.userUsecase.UpdateUser(ctx, reqModel); err != nil {
		log.Errorf("[UserController] UpdateUser - 4: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

// DeleteUser implements UserControllerInterface
func (u *userController) DeleteUser(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Params("id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User id is required",
		})
	}

	id := conv.StringToUint(userID)

	if err := u.userUsecase.DeleteUser(ctx, id); err != nil {
		log.Errorf("[UserController] DeleteUser - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

// GetUserByRoleName implements UserControllerInterface
func (u *userController) GetUserByRoleName(c fiber.Ctx) error {
	ctx := c.Context()

	roleName := c.Params("roleName")
	if roleName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Role name is required",
		})
	}

	users, err := u.userUsecase.GetUserByRoleName(ctx, roleName)
	if err != nil {
		log.Errorf("[UserController] GetUserByRoleName - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resp := []response.UserResponse{}
	for _, user := range users {
		roleNameFound := ""
		if len(user.Roles) > 0 {
			roleNameFound = user.Roles[0].Name
		}

		resp = append(resp, response.UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Photo:    user.Photo,
			Phone:    user.Phone,
			RoleName: roleNameFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Users fetched successfully",
		"data":    resp,
	})
}

// AssignUserToRole implements UserControllerInterface
func (u *userController) AssignUserToRole(c fiber.Ctx) error {
	ctx := c.Context()

	req := request.AssignUserToRoleRequest{}
	if err := c.Bind().Body(&req); err != nil {
		log.Errorf("[UserController] AssignUserToRole - 1: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[UserController] AssignUserToRole - 2: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := u.userUsecase.AssignUserToRole(ctx, req.UserID, req.RoleID); err != nil {
		log.Errorf("[UserController] AssignUserToRole - 3: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User assigned to role successfully",
	})
}

// EditAssignUserToRole implements UserControllerInterface
func (u *userController) EditAssignUserToRole(c fiber.Ctx) error {
	ctx := c.Context()
	req := request.AssignUserToRoleRequest{}

	if err := c.Bind().Body(&req); err != nil {
		log.Errorf("[UserController] EditAssignUserToRole - 1: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[UserController] EditAssignUserToRole - 2: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userRoleIDStr := c.Params("id")
	userRoleID := conv.StringToUint(userRoleIDStr)

	if err := u.userUsecase.EditAssignUserToRole(ctx, userRoleID, req.UserID, req.RoleID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user role updated successfully",
	})
}

// GetUserRoleByID implements UserControllerInterface
func (u *userController) GetUserRoleByID(c fiber.Ctx) error {
	ctx := c.Context()

	assignRoleID := c.Params("id")
	if assignRoleID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Assign role id is required",
		})
	}

	id := conv.StringToUint(assignRoleID)

	userRole, err := u.userUsecase.GetUserRoleByID(ctx, id)
	if err != nil {
		log.Errorf("[UserController] GetUserRoleByID - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resp := response.UserResponse{
		ID:       userRole.User.ID,
		Name:     userRole.User.Name,
		Email:    userRole.User.Email,
		Photo:    userRole.User.Photo,
		Phone:    userRole.User.Phone,
		RoleName: userRole.Role.Name,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User role fetched successfully",
		"data":    resp,
	})
}

// GetAllUserRoles implements UserControllerInterface
func (u *userController) GetAllUserRoles(c fiber.Ctx) error {
	ctx := c.Context()

	var req request.GetAllUserRolesRequest
	if err := c.Bind().Query(&req); err != nil {
		log.Errorf("[UserController] GetAllUserRoles - 1: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[UserController] GetAllUserRoles - 2: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}

	userRoles, total, err := u.userUsecase.GetAllUserRoles(ctx, req.Page, req.Limit, req.Search, req.SortBy, req.SortOrder)
	if err != nil {
		log.Errorf("[UserController] GetAllUserRoles - 3: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resp := []response.UserResponse{}
	for _, userRole := range userRoles {
		resp = append(resp, response.UserResponse{
			ID:       userRole.User.ID,
			Name:     userRole.User.Name,
			Email:    userRole.User.Email,
			Phone:    userRole.User.Phone,
			Photo:    userRole.User.Photo,
			RoleName: userRole.Role.Name,
		})
	}

	paginationInfo := pagination.CalculatePagination(req.Page, req.Limit, int(total))

	result := response.GetAllUsersResponse{
		Users:      resp,
		Pagination: paginationInfo,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User roles fetched successfully",
		"data":    result,
	})
}
