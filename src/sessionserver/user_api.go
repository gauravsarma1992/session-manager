package sessionserver

import (
	"errors"
	"strconv"

	"github.com/gauravsarma1992/src/sessionmgmt"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func (server *Server) GetUsersHandler(c *gin.Context) {
	var (
		users  []*sessionmgmt.User
		result *gorm.DB
	)

	if result = server.Db.Find(&users); result.Error != nil {
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"users":   users,
	})
	return
}

func (server *Server) GetUserHandler(c *gin.Context) {
	var (
		result  *gorm.DB
		reqUser *sessionmgmt.User
	)
	reqUser = &sessionmgmt.User{}
	userId := c.Param("id")
	if result = server.Db.First(reqUser, userId); result.Error != nil {
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"user":    reqUser,
	})
	return
}

func (server *Server) CreateUserHandler(c *gin.Context) {
	var (
		err    error
		result *gorm.DB
	)
	reqUser := &sessionmgmt.User{}
	if err = c.ShouldBindJSON(reqUser); err != nil {
		c.JSON(400, gin.H{
			"message": "failure",
			"error":   err.Error(),
		})
		return
	}
	if result = server.Db.Create(reqUser); result.Error != nil {
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"user":    reqUser,
	})
	return
}

func (server *Server) UpdateUserHandler(c *gin.Context) {
	var (
		result *gorm.DB
		userId int
		err    error
	)
	userIdS := c.Param("id")
	if userId, err = strconv.Atoi(userIdS); err != nil {
		c.JSON(400, gin.H{
			"message": "failure",
			"error":   err.Error(),
		})
	}
	reqUser := &sessionmgmt.User{}
	if err = c.ShouldBindJSON(reqUser); err != nil {
		c.JSON(400, gin.H{
			"message": "failure",
			"error":   err.Error(),
		})
		return
	}
	user := &sessionmgmt.User{}
	server.Db.First(user, userId)
	if user.ID == 0 {
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   errors.New("User not found"),
		})
		return
	}

	user.Email = reqUser.Email
	user.Mobile = reqUser.Password

	if result = server.Db.Save(user); result.Error != nil {
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"user":    user,
	})
	return
}

func (server *Server) DeleteUserHandler(c *gin.Context) {
	var (
		result *gorm.DB
		userId int
		err    error
	)
	userIdS := c.Param("id")
	if userId, err = strconv.Atoi(userIdS); err != nil {
		c.JSON(400, gin.H{
			"message": "failure",
			"error":   err.Error(),
		})
	}
	if result = server.Db.Delete(&sessionmgmt.User{ID: uint(userId)}); result.Error != nil {
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}
