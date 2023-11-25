package sessionserver

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gauravsarma1992/src/sessionmgmt"
)

const (
	DatabaseName = "session_users"
)

type (
	Server struct {
		config       *Config
		Db           *gorm.DB
		sessionStore *sessionmgmt.SessionStore
		apiEngine    *gin.Engine
	}
	Config struct {
		Server struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"server"`
		Db struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"db"`
	}
)

func GetDefaultConfig() (config *Config) {
	config = &Config{}
	config.Server.Host = "127.0.0.1"
	config.Server.Port = "8090"
	config.Db.Host = "127.0.0.1"
	config.Db.Port = "3306"
	config.Db.Username = "root"
	config.Db.Password = ""
	return
}

func NewServer(config *Config) (server *Server, err error) {
	if config == nil {
		config = GetDefaultConfig()
	}

	server = &Server{
		config:    config,
		apiEngine: gin.Default(),
	}
	server.setupRoutes()
	server.setupDb()
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
	server.apiEngine.PUT("/api/users/:id", server.UpdateUserHandler)
	server.apiEngine.DELETE("/api/users/:id", server.DeleteUserHandler)
	return
}

func (server *Server) setupDb() (err error) {
	var (
		db *gorm.DB
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		server.config.Db.Username,
		server.config.Db.Password,
		server.config.Db.Host,
		server.config.Db.Port,
		DatabaseName,
	)
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	}
	server.Db = db
	server.Db.AutoMigrate(&sessionmgmt.User{})
	return
}

func (server *Server) Run() (err error) {
	serverAddress := fmt.Sprintf("%s:%s",
		server.config.Server.Host,
		server.config.Server.Port,
	)
	log.Println("Starting server on", serverAddress)
	err = server.apiEngine.Run(serverAddress)
	return
}
