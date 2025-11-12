package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	//CreateToken is an function for creation a token for scpecific username and duration
	CreateToken(username string, role string, duration time.Duration, tokenType TokenType) (string, error)
	//VerifyToken checks if input token is valid or not
	VerifyToken(token string, tokenTypy TokenType) (*Payload, error)
}
