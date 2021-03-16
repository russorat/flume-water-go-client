package flumewater

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestFlumeFetchUserDevices(t *testing.T) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := NewClient(os.Getenv("FLUME_CLIENT_ID"), os.Getenv("FLUME_CLIENT_SECRET"), os.Getenv("FLUME_USERNAME"), os.Getenv("FLUME_PASSWORD"))
	user, err := client.FetchUserDevices()
	assert.Nil(t, err)
	assert.Equal(t, client.userID, user.Data[0].UserID)
}

func TestFlumeFetchUserDevice(t *testing.T) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := NewClient(os.Getenv("FLUME_CLIENT_ID"), os.Getenv("FLUME_CLIENT_SECRET"), os.Getenv("FLUME_USERNAME"), os.Getenv("FLUME_PASSWORD"))
	devices, _ := client.FetchUserDevices()
	device, err := client.FetchUserDevice(devices.Data[0].ID)
	assert.Nil(t, err)
	assert.Equal(t, client.userID, device.Data[0].UserID)
	assert.Equal(t, devices.Data[0].ID, device.Data[0].ID)
}
