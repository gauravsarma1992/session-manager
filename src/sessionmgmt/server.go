package sessionmgmt

import (
	"fmt"

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
		config = &Config{}
		config.Host = "127.0.0.1"
		config.Port = "9090"
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
	server.apiEngine.GET("/sessions/:id", server.GetSessionHandler)
	server.apiEngine.POST("/sessions", server.CreateSessionHandler)
	server.apiEngine.DELETE("/sessions/:id", server.DeleteSessionHandler)
	return
}

func (server *Server) Run() (err error) {
	err = server.apiEngine.Run(
		fmt.Sprintf("%s:%s",
			server.config.Host,
			server.config.Port,
		),
	)
	return
}
