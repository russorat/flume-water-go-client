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
	devices, err := client.FetchUserDevices(FlumeWaterFetchDeviceRequest{IncludeUser: true, IncludeLocation: true})
	assert.Nil(t, err)
	assert.Equal(t, client.userID, devices[0].UserID)
}

func TestFlumeFetchUserDevice(t *testing.T) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := NewClient(os.Getenv("FLUME_CLIENT_ID"), os.Getenv("FLUME_CLIENT_SECRET"), os.Getenv("FLUME_USERNAME"), os.Getenv("FLUME_PASSWORD"))
	devices, _ := client.FetchUserDevices(FlumeWaterFetchDeviceRequest{})
	device, err := client.FetchUserDevice(devices[0].ID, FlumeWaterFetchDeviceRequest{IncludeUser: true, IncludeLocation: true})
	assert.Nil(t, err)
	assert.Equal(t, client.userID, device.UserID)
	assert.Equal(t, devices[0].ID, device.ID)
}
