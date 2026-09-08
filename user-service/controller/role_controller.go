package controller

import (
	"pos/user-service/controller/request"
	"pos/user-service/controller/response"
	"pos/user-service/model"
	"pos/user-service/pkg/conv"
	"pos/user-service/pkg/validator"
	"pos/user-service/usecase"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type RoleControllerInterface interface {
	CreateRole(c fiber.Ctx) error
	UpdateRole(c fiber.Ctx) error
	DeleteRole(c fiber.Ctx) error
	GetRoleByID(c fiber.Ctx) error
	GetAllRoles(c fiber.Ctx) error
}

type roleController struct {
	roleUsecase usecase.RoleUseCaseInterface
}

func NewRoleController(roleUsecase usecase.RoleUseCaseInterface) RoleControllerInterface {
	return &roleController{roleUsecase: roleUsecase}
}

// CreateRole implements RoleControllerInterface
func (r *roleController) CreateRole(c fiber.Ctx) error {
	ctx := c.Context()

	req := request.CreateRoleRequest{}

	if err := c.Bind().Body(&req); err != nil {
		log.Errorf("[RoleController] CreateRole - 1: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[RoleController] CreateRole - 2: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reqModel := model.Role {
		Name: req.Name,
	}

	if err := r.roleUsecase.CreateRole(ctx, reqModel); err != nil {
		log.Errorf("[RoleController] CreateRole - 3: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Role created successfully",
	})
}

// UpdateRole implements RoleControllerInterface
func (r *roleController) UpdateRole(c fiber.Ctx) error {
	ctx := c.Context()

	req := request.CreateRoleRequest{}
	if err := c.Bind().Body(&req); err != nil {
		log.Errorf("[RoleController] UpdateRole - 1: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[RoleController] UpdateRole - 2: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reqModel := model.Role{
		ID: conv.StringToUint(c.Params("id")),
		Name: req.Name,
	}

	if err := r.roleUsecase.UpdateRole(ctx, reqModel); err != nil {
		log.Errorf("[RoleController] UpdateRole - 3: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Role updated successfully",
	})
}

// DeleteRole implements RoleControllerInterface
func (r *roleController) DeleteRole(c fiber.Ctx) error {
	ctx := c.Context()

	roleID := c.Params("id")
	if roleID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Role ID is required",
		})
	}

	id := conv.StringToUint(roleID)

	if err := r.roleUsecase.DeleteRole(ctx, id); err != nil {
		log.Errorf("[RoleController] DeleteRole - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Role delete successfully",
	})
}

// GetRoleByID implements RoleControllerInterface
func (r *roleController) GetRoleByID(c fiber.Ctx) error {
	ctx := c.Context()

	roleID := c.Params("id")
	if roleID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Role id is required",
		})
	}

	id := conv.StringToUint(roleID)

	role, err := r.roleUsecase.GetRoleByID(ctx, id)
	if err != nil {
		log.Errorf("[RoleController] GetRoleByID - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "role fatched data successfuly",
		"data": role,
	})
}

// GetAllRoles implements RoleControllerInterface
func (r *roleController) GetAllRoles(c fiber.Ctx) error {
	ctx := c.Context()

	roles, totalRecord, err := r.roleUsecase.GetAllRoles(ctx)
	if err != nil {
		log.Errorf("[RoleController] GetAllRole - 1: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"total": totalRecord,
			"message": err.Error(),
		})

	}

	resp := []response.RoleResponse{}
	for _, role := range roles {
		resp = append(resp, response.RoleResponse{
			ID: role.ID,
			Name: role.Name,
			CountUsers: int64(len(role.Users)),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Role fetched successfully",
		"data": resp,
	})
}