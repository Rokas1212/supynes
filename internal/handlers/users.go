package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  struct {
		Email       string `json:"email"`
		DisplayName string `json:"displayName"`
	} `json:"user"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// TODO: ADD VALIDATION (anything is allowed in for now)

	resp := LoginResponse{
		Token: "dummy-wummy",
		User: struct {
			Email       string `json:"email"`
			DisplayName string `json:"displayName"`
		}{
			Email:       req.Email,
			DisplayName: "User",
		},
	}
	c.JSON(http.StatusOK, resp)
}
