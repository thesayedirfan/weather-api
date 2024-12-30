package http

import "github.com/gin-gonic/gin"

func (h *HttpHandler) Health(c *gin.Context) {
	c.Status(200)
	c.Abort()
	return
}
