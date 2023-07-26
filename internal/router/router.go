package router

import (
	"github.com/gin-gonic/gin"
	"github.com/onewesong/goforward/internal/api/forward"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()

	g.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})
	apiForward := g.Group("/api/forward")
	apiForward.GET("", forward.GetAll)
	apiForward.GET("/:listen_addr", forward.Get)
	apiForward.POST("", forward.Add)
	apiForward.DELETE("/:listen_addr", forward.Del)
	return g
}
