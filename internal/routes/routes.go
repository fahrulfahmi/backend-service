package routes

import (
	"backend-service/internal/auth"
	"backend-service/internal/category"
	"backend-service/internal/menu"
	"backend-service/internal/middleware"
	"backend-service/internal/order"
	"backend-service/internal/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	authHandler := auth.NewHandler(db)
	userHandler := user.NewHandler(db)
	categoryHandler := category.NewHandler(db)
	menuHandler := menu.NewHandler(db)
	orderHandler := order.NewHandler(db)

	r.POST("/auth/login", authHandler.Login)

	owner := r.Group("/owner")
	owner.Use(middleware.JWT(), middleware.RequireRole("OWNER"))
	{
		owner.POST("/users", userHandler.Create)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.JWT(), middleware.RequireRole("OWNER", "ADMIN"))
	{
		admin.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "admin ok"})
		})

		admin.POST("/categories", categoryHandler.Create)
		admin.POST("/menus", menuHandler.Create)
	}

	cashier := r.Group("/cashier")
	cashier.Use(middleware.JWT(), middleware.RequireRole("CASHIER"))
	{
		cashier.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "cashier ok"})
		})
	}

	r.GET("/categories", categoryHandler.List)
	r.GET("/menus", menuHandler.List)

	r.POST("/orders", orderHandler.Create)
	r.GET("/orders/:id", orderHandler.Get)
}
