package flumewater

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestFlumeFetchUserLocations(t *testing.T) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := NewClient(os.Getenv("FLUME_CLIENT_ID"), os.Getenv("FLUME_CLIENT_SECRET"), os.Getenv("FLUME_USERNAME"), os.Getenv("FLUME_PASSWORD"))
	var queryParams = NewLocationsQueryParams()
	queryParams.Limit = 10
	locations, err := client.FetchUserLocations(*queryParams)
	assert.Nil(t, err)
	assert.Equal(t, 1, locations.Count)
}

func TestFlumeFetchUserLocation(t *testing.T) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := NewClient(os.Getenv("FLUME_CLIENT_ID"), os.Getenv("FLUME_CLIENT_SECRET"), os.Getenv("FLUME_USERNAME"), os.Getenv("FLUME_PASSWORD"))
	var queryParams = NewLocationsQueryParams()
	queryParams.Limit = 10
	locations, _ := client.FetchUserLocations(*queryParams)
	location, err := client.FetchUserLocation(locations.Data[0].ID)
	assert.Nil(t, err)
	assert.Equal(t, 1, location.Count)
}
