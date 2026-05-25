package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ai-project-os/api/internal/config"
	"github.com/ai-project-os/api/internal/database"
	"github.com/ai-project-os/api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	cfg := config.Load()

	if err := database.Init(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	email := getArg(1, "admin@aiproject.os")
	username := getArg(2, "admin")
	password := getArg(3, "admin123")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	var user models.User
	result := database.GetDB().Where("email = ?", email).First(&user)

	if result.Error != nil {
		// Not found — create new
		user = models.User{
			Email:        email,
			Username:     username,
			PasswordHash: string(hash),
			Role:         models.RoleAdmin,
			Status:       models.UserStatusActive,
		}
		if err := database.GetDB().Create(&user).Error; err != nil {
			log.Fatalf("Failed to create admin user: %v", err)
		}
		fmt.Printf("✅ Admin user created successfully!\n")
	} else {
		// Already exists — update role, password, status
		user.Role = models.RoleAdmin
		user.Status = models.UserStatusActive
		user.PasswordHash = string(hash)
		if err := database.GetDB().Save(&user).Error; err != nil {
			log.Fatalf("Failed to update admin user: %v", err)
		}
		fmt.Printf("✅ Admin user already exists — updated role and password.\n")
	}
	fmt.Printf("   Email:    %s\n", email)
	fmt.Printf("   Username: %s\n", username)
	fmt.Printf("   Password: %s\n", password)
	fmt.Printf("   Role:     admin\n")
	fmt.Printf("\n   Login at: /login\n")
}

func getArg(index int, fallback string) string {
	if len(os.Args) > index && os.Args[index] != "" {
		return os.Args[index]
	}
	return fallback
}
