package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"backend-service/internal/category"
	"backend-service/internal/config"
	"backend-service/internal/menu"
	"backend-service/internal/order"
	"backend-service/internal/routes"
	"backend-service/internal/user"
)

func main() {
	config.LoadEnv()

	db := config.InitDatabase()

	// 🔥 AUTO MIGRATE SEMUA MODEL (WAJIB)
	db.AutoMigrate(
		&user.User{},
		&category.Category{},
		&menu.Menu{},
		&order.Order{},
		&order.OrderItem{},
	)

	r := gin.Default()

	routes.RegisterRoutes(r, db)

	log.Println("🚀 Server running on :8080")
	r.Run(":8080")
}
