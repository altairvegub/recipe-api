package ports

import (
	"errors"
	"net/http"
	"recipe/internal/service"

	"github.com/gin-gonic/gin"
)

func HandleHealth() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Status(200)
	}
}

func HandleSignup(svc *service.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		var signupRequest SignupRequest
		if err := c.ShouldBindJSON(&signupRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := (*svc).Signup(signupRequest.Email, signupRequest.Password)

		if err != nil {
			if errors.Is(err, service.ErrResourceAlreadyExists) {
				c.Status(http.StatusConflict)
				return
			}
			c.Status(http.StatusInternalServerError)
			return
		}

		//c.Status(http.StatusCreated)
		c.JSON(http.StatusCreated, token)
	}
}
