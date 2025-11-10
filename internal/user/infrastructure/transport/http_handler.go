package transport

import (
	"github/gjangra9988/go-ddd/internal/user/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *application.UserRepository
}

func NewHandler(service *application.UserRepository) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	userRoute := r.Group("/user")

	userRoute.POST("/", h.createUser)
	userRoute.GET("/:id", h.getUser)
	userRoute.PUT("/:id", h.updateUser)
	userRoute.DELETE("/:id", h.deleteUser)
}

func (h *Handler) createUser(c *gin.Context){
	
	var req application.UserCreateRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid Data"})
		return
	}
	id, err := h.service.CreateUser(c.Request.Context(), req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": id})
}

func (h *Handler) getUser(c *gin.Context){
	
	id := c.Param("id")

	user, err := h.service.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusFound, gin.H{"message": user})
}

func (h *Handler) updateUser(c *gin.Context){

	id := c.Param("id")

	var req application.UserUpdateRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data"})
		return
	}

	user, err := h.service.UpdateUser(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Internal Server Error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User detail updated", "user": user})
}

func (h *Handler) deleteUser(c *gin.Context){

	id := c.Param("id")

	err := h.service.DeleteUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message" : "User deleted"})
}

