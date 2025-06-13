package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiuncy/spp/repository"
)

type userHandler struct {
	userRepo repository.UserRepo
}

func NewUserHandler(userRepo repository.UserRepo) *userHandler {
	return &userHandler{userRepo}
}

func (u *userHandler) CreateUser(c *gin.Context) {
	panic("unimplemented")
}

func (u *userHandler) Login(c *gin.Context) {
	type userLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var login userLogin

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := u.userRepo.Login(login.Email, login.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
