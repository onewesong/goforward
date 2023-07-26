package forward

import (
	"github.com/gin-gonic/gin"
	"github.com/onewesong/goforward/internal/manager"
)

// /api/forward [GET]
// 获取当前的所有转发信息
func GetAll(c *gin.Context) {
	m := manager.GetInstance()
	c.JSON(200, m.ForwardMap)
}

// /api/forward/:listen_addr [GET]
// 获取指定监听地址的转发信息
func Get(c *gin.Context) {
	m := manager.GetInstance()
	c.JSON(200, m.ForwardMap[c.Param("listen_addr")])
}
