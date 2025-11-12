package token

import (
	"mySimpleBank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwnerName()
	duration := time.Minute
	role := util.RandomString(6)

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)
	tokenType := TokenType(TokenTypeAccessToken)

	token, err := maker.CreateToken(username, role, duration, tokenType)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token, tokenType)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}
