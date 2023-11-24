package sessionserver

import (
	"github.com/gin-gonic/gin"

	"github.com/gauravsarma1992/src/sessionmgmt"
)

func (server *Server) GetSessionHandler(c *gin.Context) {
	sessionId := c.Param("id")
	session, err := server.sessionStore.GetSession(sessionmgmt.TokenT(sessionId))
	if err != nil {
		c.JSON(401, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":   "success",
		"session":   session,
		"is_active": session.IsActive(),
	})
	return
}

func (server *Server) CreateSessionHandler(c *gin.Context) {
	var (
		err error
	)
	session := &sessionmgmt.Session{}
	if err = c.ShouldBindJSON(session); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	if session, err = server.sessionStore.AddSession(session.Token, session.SessionObj); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

func (server *Server) DeleteSessionHandler(c *gin.Context) {
	sessionId := c.Param("id")
	err := server.sessionStore.RemoveSession(sessionmgmt.TokenT(sessionId))
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}
