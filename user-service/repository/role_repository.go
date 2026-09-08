package repository

// ROLE: create, update, delete dan get

import (
	"context"
	"errors"
	"pos/user-service/model"

	"github.com/gofiber/fiber/v3/log"
	"gorm.io/gorm"
)

type RoleRepositoryInterface interface {
	CreateRole(ctx context.Context, role model.Role) error
	UpdateRole(ctx context.Context, role model.Role) error
	DeleteRole(ctx context.Context, id uint) error
	GetRoleByID(ctx context.Context, id uint) (*model.Role, error)
	GetAllRoles(ctx context.Context) ([]model.Role, int64, error)
	
}

type roleRepository struct {
	db *gorm.DB
}

// CreateRole implements RoleRepositoryInterface
func (r *roleRepository) CreateRole(ctx context.Context, role model.Role) error {
	select {
	case <-ctx.Done():
		log.Errorf("[RoleRepository] CreateRole - 1: %v", ctx.Err())
		return ctx.Err()

	default:
		return r.db.WithContext(ctx).Create(&role).Error
	}
}

// DeleteRole implements RoleRepositoryInterface
func (r *roleRepository) DeleteRole(ctx context.Context, id uint) error {
	select {
	case <-ctx.Done():
		log.Errorf("[RoleRepository] DeleteRole - 1: %v", ctx.Err())
		return ctx.Err()

	default:
		modelRole := model.Role{}
		if err := r.db.WithContext(ctx).Preload("Users").Where("id = ?", id).First(&modelRole).Error; err != nil {
			log.Errorf("[RoleRepository] DeleteRole - 2: %v", err)
			return err
		}

		if modelRole.ID == 0 {
			log.Errorf("[RoleRepository] DeleteRole - 3: Role with ID %v", "Role not found")
			return errors.New("role not found")
		}

		if len(modelRole.Users) > 0 {
			log.Errorf("[RoleRepository] DeleteRole - 4: %v", "Role has users")
			return errors.New("role has users")
		}

		return r.db.WithContext(ctx).Delete(&model.Role{}, id).Error
	}
}

// GetAllRoles implements RoleRepositoryInterface
func (r *roleRepository) GetAllRoles(ctx context.Context) ([]model.Role, int64, error) {
	select {
	case <-ctx.Done():
		log.Errorf("[RoleRepository] GetAllRoles - 1: %v", ctx.Err())
		return nil, 0, ctx.Err()

	default:
		modelRoles := []model.Role{}
		err := r.db.WithContext(ctx).Preload("Users").Find(&modelRoles).Error
		if err != nil {
			log.Errorf("[RoleRepository] GetAllRoles - 2: %v", err)
			return nil, 0, err
		}

		return modelRoles, int64(len(modelRoles)), nil
	}
}

// GetRoleByID implements RoleRepositoryInterface
func (r *roleRepository) GetRoleByID(ctx context.Context, id uint) (*model.Role, error) {
	select {
	case <-ctx.Done():
		log.Errorf("[RoleRepository] GetRoleByID - 1: %v", ctx.Err())
		return nil, ctx.Err()

	default:
		modelRole := model.Role{}
		if err := r.db.WithContext(ctx).Preload("Users").Where("id = ?", id).First(&modelRole).Error; err != nil {
			log.Errorf("[RoleRepository] GetRoleByID - 2: %v", err)
			return nil, err
		}
		return &modelRole, nil
	}
}

// UpdateRole implements RoleRepositoryInterface
func (r *roleRepository) UpdateRole(ctx context.Context, role model.Role) error {
	select {
	case <-ctx.Done():
		log.Errorf("[RoleRepository] UpdateRole - 1: %v", ctx.Err())
		return ctx.Err()

	default:
		modelRole := model.Role{}
		if err := r.db.WithContext(ctx).Where("id = ?", role.ID).First(&modelRole).Error; err != nil {
			log.Errorf("[RoleRepository] UpdateRole - 2: %v", err)
			return err
		}

		modelRole.Name = role.Name
		return r.db.WithContext(ctx).Save(&modelRole).Error
	}
}

func NewRoleRepository(db *gorm.DB) RoleRepositoryInterface {
	return &roleRepository{db: db}
}
