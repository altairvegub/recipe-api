package ports

import (
	"errors"
	"fmt"
	"net/http"

	"recipe/pkg/log"

	"github.com/gin-gonic/gin"
)

func RunHTTPServer(port int, logger log.Logger) *http.Server {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
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
