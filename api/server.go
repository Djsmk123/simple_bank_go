package api

import (
	db "github.com/djsmk123/simplebank/db/sqlc"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/create-account", server.createAccount)
	router.GET("/", index)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/create-transfer", server.createTransfer)

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
