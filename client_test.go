package flumewater

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestFlumeCreateClient(t *testing.T) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := NewClient(os.Getenv("FLUME_CLIENT_ID"), os.Getenv("FLUME_CLIENT_SECRET"), os.Getenv("FLUME_USERNAME"), os.Getenv("FLUME_PASSWORD"))
	assert.Equal(t, os.Getenv("FLUME_CLIENT_ID"), client.ClientID)
	assert.Equal(t, os.Getenv("FLUME_CLIENT_SECRET"), client.ClientSecret)
	assert.Equal(t, os.Getenv("FLUME_USERNAME"), client.Username)
	assert.Equal(t, os.Getenv("FLUME_PASSWORD"), client.Password)
}
