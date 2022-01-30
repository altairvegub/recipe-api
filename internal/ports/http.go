package ports

import (
	"errors"
	"fmt"
	"net/http"

	"recipe/pkg/log"

	"github.com/gin-gonic/gin"

	"recipe/internal/service"
)

func RunHTTPServer(port int, logger log.Logger, svc service.Service) *http.Server {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})
	r.POST("/signup", func(c *gin.Context) {
		var signupRequest SignupRequest
		if err := c.ShouldBindJSON(&signupRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := svc.Signup(signupRequest.Email, signupRequest.Password)
		if err != nil {
			if errors.Is(err, service.ErrResourceAlreadyExists) {
				c.Status(http.StatusConflict)
				return
			}
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusCreated)
	})

	httpServerPort := fmt.Sprintf(":%d", port)
	r.Run(httpServerPort)

	srv := &http.Server{
		Addr:    httpServerPort,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logger.Infof("server is running at %v", httpServerPort)
		}
	}()

	return srv
}
