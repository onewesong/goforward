package forward

import (
	"github.com/gin-gonic/gin"
	"github.com/onewesong/goforward/internal/manager"
	"github.com/onewesong/goforward/internal/models"
)

type AddForwardRequest struct {
	ForwardLink string `json:"forward_link"`
	Override    bool   `json:"override"`
}

// /api/forward [POST]
// 添加转发
func Add(c *gin.Context) {
	var request AddForwardRequest
	c.ShouldBindJSON(&request)
	if len(request.ForwardLink) == 0 {
		c.JSON(400, gin.H{"error": "forward link is empty"})
		return
	}
	forwardLinks, err := models.NewForwardLinks(request.ForwardLink)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	m := manager.GetInstance()
	err = m.AddForward(forwardLinks, request.Override)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "ok")
}
