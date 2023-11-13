package sessionmgmt

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) GetSessionHandler(c *gin.Context) {
	sessionId := c.Param("id")
	session, err := server.sessionStore.GetSession(TokenT(sessionId))
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"session": session,
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
