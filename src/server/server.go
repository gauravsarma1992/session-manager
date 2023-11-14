package sessionmgmt

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/gauravsarma1992/src/sessionmgmt"
)

type (
	Server struct {
		config       *Config
		sessionStore *sessionmgmt.SessionStore
		apiEngine    *gin.Engine
	}
	Config struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
)

func NewServer(config *Config) (server *Server, err error) {
	if config == nil {
		config = &Config{
			Host: "127.0.0.1",
			Port: "8090",
		}
	}

	server = &Server{
		config:    config,
		apiEngine: gin.Default(),
	}
	server.setupRoutes()
	server.sessionStore, _ = sessionmgmt.NewSessionStore(nil)
	return
}

func (server *Server) setupRoutes() (err error) {
	server.apiEngine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.apiEngine.GET("/api/sessions/:id", server.GetSessionHandler)
	server.apiEngine.POST("/api/sessions", server.CreateSessionHandler)
	server.apiEngine.DELETE("/api/sessions/:id", server.DeleteSessionHandler)

	server.apiEngine.GET("/api/users", server.GetUsersHandler)
	server.apiEngine.GET("/api/users/:id", server.GetUserHandler)
	server.apiEngine.POST("/api/users", server.CreateUserHandler)
	server.apiEngine.PUT("/api/users", server.UpdateUserHandler)
	server.apiEngine.DELETE("/api/users/:id", server.DeleteUserHandler)
	return
}

func (server *Server) Run() (err error) {
	serverAddress := fmt.Sprintf("%s:%s",
		server.config.Host,
		server.config.Port,
	)
	log.Println("Starting server on", serverAddress)
	err = server.apiEngine.Run(serverAddress)
	return
}
