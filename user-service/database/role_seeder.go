package database

import (
	"pos/user-service/model"

	"github.com/gofiber/fiber/v3/log"
	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) {
	roles := []model.Role{
		{Name: "Manager"},
		{Name: "Keeper"},
	}

	for _, role := range roles {
		if err := db.Where(model.Role{Name: role.Name}).FirstOrCreate(&role).Error; err != nil {
			log.Errorf("[RoleSeeder] SeedRole - 1: %v", err)
		} else {
			log.Infof("[RoleSeeder] SeedRole - 2: %v", "Role seeded successfully")
		}
	}
}