package flumewater

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestFlumeGetToken(t *testing.T) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := NewClient(os.Getenv("FLUME_CLIENT_ID"), os.Getenv("FLUME_CLIENT_SECRET"), os.Getenv("FLUME_USERNAME"), os.Getenv("FLUME_PASSWORD"))
	client.GetToken()
	assert.NotEmpty(t, client.AuthData)
	assert.NotEmpty(t, client.AuthData.AccessToken)
	assert.Equal(t, 604800, client.AuthData.ExpiresIn)
	assert.NotEmpty(t, client.AuthData.RefreshToken)
	assert.Equal(t, "bearer", client.AuthData.TokenType)
	assert.Equal(t, 29571, client.userID)
}

func TestFlumeRefreshToken(t *testing.T) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := NewClient(os.Getenv("FLUME_CLIENT_ID"), os.Getenv("FLUME_CLIENT_SECRET"), os.Getenv("FLUME_USERNAME"), os.Getenv("FLUME_PASSWORD"))
	client.GetToken()
	client.RefreshToken()
	assert.NotEmpty(t, client.AuthData)
	assert.NotEmpty(t, client.AuthData.AccessToken)
	assert.Equal(t, 604800, client.AuthData.ExpiresIn)
	assert.NotEmpty(t, client.AuthData.RefreshToken)
	assert.Equal(t, "bearer", client.AuthData.TokenType)
	assert.Equal(t, 29571, client.userID)
}
