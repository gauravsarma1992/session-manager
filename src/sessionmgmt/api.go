package sessionmgmt

import "github.com/gin-gonic/gin"

func (server *Server) GetSessionHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

func (server *Server) CreateSessionHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

func (server *Server) DeleteSessionHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}
