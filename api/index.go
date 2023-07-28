package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexRequest struct {
	Status    int  `json:"status"`
	Listening bool `json:"listening"`
}

func index(ctx *gin.Context) {
	var req IndexRequest = IndexRequest{
		Status:    200,
		Listening: true,
	}
	ctx.JSON(http.StatusOK, req)
}
