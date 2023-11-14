package sessionmgmt

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) GetUsersHandler(c *gin.Context) {
	//userId := c.Param("id")
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

func (server *Server) GetUserHandler(c *gin.Context) {
	//userId := c.Param("id")
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

func (server *Server) CreateUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

func (server *Server) UpdateUserHandler(c *gin.Context) {
	//userId := c.Param("id")
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

func (server *Server) DeleteUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}
