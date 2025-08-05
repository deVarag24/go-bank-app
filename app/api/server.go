package api

import (
	db "my-go-app/app/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct{
	store *db.Store
	router *gin.Engine
}

// create a new http server and setup routing
func NewServer(store *db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()

	// define routes
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccountById)
	router.GET("/accounts", server.getAllAccount)


	server.router = router
	return server
}

// start HTTP server
func (server *Server) StartServer(address string) error{
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}