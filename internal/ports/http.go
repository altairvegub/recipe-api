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
	r.GET("/health", HandleHealth())
	r.POST("/signup", HandleSignup(&svc))

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
