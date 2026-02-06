package menu

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	service *Service
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{service: NewService(db)}
}

type CreateMenuRequest struct {
	Name       string `json:"name"`
	Price      int64  `json:"price"`
	CategoryID uint   `json:"category_id"`
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(req.Name, req.Price, req.CategoryID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "menu created"})
}

func (h *Handler) List(c *gin.Context) {
	menus, _ := h.service.List()
	c.JSON(http.StatusOK, menus)
}
