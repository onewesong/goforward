package forward

import (
	"github.com/gin-gonic/gin"
	"github.com/onewesong/goforward/internal/manager"
	"github.com/onewesong/goforward/internal/models"
)

// /api/forward/:listen_addr [DELETE]
// 获取指定监听地址的转发信息
func Del(c *gin.Context) {
	m := manager.GetInstance()
	listenAddr := c.Param("listen_addr")
	if listenAddr == "" {
		c.JSON(400, gin.H{"error": "listen_addr is empty"})
		return
	}
	if _, err := models.ParseAddr(listenAddr); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := m.DelForward(listenAddr)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "ok")
}
