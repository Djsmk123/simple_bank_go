package api

import (
	"os"
	"testing"
	"time"

	db "github.com/djsmk123/simplebank/db/sqlc"
	"github.com/djsmk123/simplebank/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokkenStructureKey:  utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
