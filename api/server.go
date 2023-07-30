package api

import (
	db "github.com/djsmk123/simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

//To create a new server

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/createAccount", server.createAccount)
	router.GET("/", index)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

//function error handler

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
