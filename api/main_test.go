package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	db "github.com/thanhqt2002/hackathon/db/sqlc"
	"github.com/thanhqt2002/hackathon/db/util"
)

func NewTestServer(t *testing.T, store *db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomStr(32),
		TokenAccessDuration: time.Minute,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}
