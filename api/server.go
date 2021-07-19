package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/thanhqt2002/hackathon/db/sqlc"
	"github.com/thanhqt2002/hackathon/db/util"
	"github.com/thanhqt2002/hackathon/token"
)

type Server struct {
	config     util.Config
	tokenMaker token.Maker
	store      *db.Store
	router     *gin.Engine
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize token maker")
	}

	server := &Server{config: config, tokenMaker: tokenMaker, store: store}
	router := gin.Default()
	router.Use(CORS())
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/users", server.CreateUser)
	router.POST("/users/login", server.LoginUser)

	authRoutes.POST("/tasks/add", server.CreateTask)
	authRoutes.GET("/tasks", server.ListTasks)
	authRoutes.GET("/tasks/admin", server.ListTasksAdmin)

	authRoutes.POST("/submissions", server.CreateSubmission)
	authRoutes.GET("/ranking", server.ListScores)

	server.router = router
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
