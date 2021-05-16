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

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize token maker")
	}

	server := &Server{config: config, tokenMaker: tokenMaker, store: store}
	router := gin.Default()
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/users", server.CreateUser)
	router.POST("/users/login", server.LoginUser)

	authRoutes.POST("/tasks", server.CreateTask)
	authRoutes.GET("/tasks/", server.ListTasks)
	authRoutes.GET("/tasks/admin", server.ListTasksAdmin)

	authRoutes.POST("/submissions", server.CreateSubmission)

	server.router = router
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
