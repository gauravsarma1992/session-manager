package sessionmgmt

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type (
	Server struct {
		config       *Config
		sessionStore *SessionStore
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
	server.sessionStore, _ = NewSessionStore(nil)
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
