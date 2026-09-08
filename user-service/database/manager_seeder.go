package database

import (
	"pos/user-service/model"
	"pos/user-service/pkg/conv"

	"github.com/gofiber/fiber/v3/log"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const defaultManagerEmail = "manager@gmail.com"

func SeedManager(db *gorm.DB) {
	// existing manager already seeded, nothing to do
	var count int64
	if err := db.Model(&model.User{}).Where("email = ?", defaultManagerEmail).Count(&count).Error; err != nil {
		log.Fatalf("%s: %v", err.Error(), err)
	}
	if count > 0 {
		log.Infof("Admin Manager already seeded, skipping")
		return
	}

	password := viper.GetString("MANAGER_DEFAULT_PASSWORD")
	if password == "" {
		if viper.GetString("APP_ENV") == "production" {
			log.Fatalf("MANAGER_DEFAULT_PASSWORD must be set in production, refusing to seed with a default password")
		}
		password = "manager123"
		log.Infof("MANAGER_DEFAULT_PASSWORD not set, falling back to default seeder password (development only)")
	}

	bytes, err := conv.HashPassword(password)
	if err != nil {
		log.Fatalf("%s: %v", err.Error(), err)
	}

	modelRole := model.Role{}
	if err := db.Where("name = ?", "Manager").First(&modelRole).Error; err != nil {
		log.Fatalf("%s: %v", err.Error(), err)
	}

	admin := model.User{
		Name:     "Manager",
		Email:    defaultManagerEmail,
		Password: bytes,
		Roles:    []model.Role{modelRole},
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("%s: %v", err.Error(), err)
	}

	log.Infof("Admin %s created", admin.Name)
}