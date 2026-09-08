package usecase

import (
	"context"
	"pos/user-service/model"
	"pos/user-service/repository"
)

type RoleUseCaseInterface interface {
	CreateRole(ctx context.Context, role model.Role) error
	UpdateRole(ctx context.Context, role model.Role) error
	DeleteRole(ctx context.Context, id uint) error
	GetRoleByID(ctx context.Context, id uint) (*model.Role, error)
	GetAllRoles(ctx context.Context) ([]model.Role, int64, error)
}

type roleUseCase struct {
	roleRepo repository.RoleRepositoryInterface
}

// CreateRole implements RoleUseCaseInterface
func (r *roleUseCase) CreateRole(ctx context.Context, role model.Role) error {
	return r.roleRepo.CreateRole(ctx, role)
}

// DeleteRole implements RoleUseCaseInterface
func (r *roleUseCase) DeleteRole(ctx context.Context, id uint) error {
	return r.roleRepo.DeleteRole(ctx, id)
}

// GetAllRoles implements RoleUseCaseInterface
func (r *roleUseCase) GetAllRoles(ctx context.Context) ([]model.Role, int64, error) {
	return r.roleRepo.GetAllRoles(ctx)
}

// GetRoleByID implements RoleUseCaseInterface
func (r *roleUseCase) GetRoleByID(ctx context.Context, id uint) (*model.Role, error) {
	return r.roleRepo.GetRoleByID(ctx, id)
}

// UpdateRole implements RoleUseCaseInterface
func (r *roleUseCase) UpdateRole(ctx context.Context, role model.Role) error {
	return r.roleRepo.UpdateRole(ctx, role)
}

func NewRoleUseCase(roleRepo repository.RoleRepositoryInterface) RoleUseCaseInterface {
	return &roleUseCase{roleRepo: roleRepo}
}