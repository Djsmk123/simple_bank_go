package api

import (
	"fmt"

	db "github.com/djsmk123/simplebank/db/sqlc"
	"github.com/djsmk123/simplebank/token"
	"github.com/djsmk123/simplebank/utils"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/lib/pq"
)

type Server struct {
	config     utils.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

//To create a new server

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPastoMaker(config.TokkenStructureKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}
func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/create-user", server.createUser)
	router.POST("/login-user", server.LoginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/user/:id", server.getUser)

	authRoutes.POST("/create-account", server.createAccount)

	router.GET("/", index)
	authRoutes.GET("/account/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.POST("/create-transfer", server.createTransfer)
	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

//function error handler

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
