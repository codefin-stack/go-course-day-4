package container

import (
	"hexagonal/internal/adapters/db"
	"hexagonal/internal/adapters/rest"
	"hexagonal/internal/application"
	"hexagonal/internal/domain"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Service struct{}

func (s *Service) Start() error {
	dbConn, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	dbConn.AutoMigrate(&domain.User{})
	// Create Repository
	userRepo := db.NewDBUserRepository(dbConn)
	// Initialize user service
	userService1 := application.NewUserService(userRepo)
	// Initialize user handler
	userHandler := rest.NewUserHandler(userService1)
	r := gin.Default()

	r.GET("/user/:id", userHandler.GetByID)
	r.POST("/users", userHandler.CreateUser)
	log.Println("Server started at :8080")
	if err := r.Run(":8080"); err != nil {
		return err
	}
	return nil
}

func NewService() *Service {
	return &Service{}
}
