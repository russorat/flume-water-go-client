package flumewater

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestFlumeQueryDevice(t *testing.T) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := NewClient(os.Getenv("FLUME_CLIENT_ID"), os.Getenv("FLUME_CLIENT_SECRET"), os.Getenv("FLUME_USERNAME"), os.Getenv("FLUME_PASSWORD"))
	devices, _ := client.FetchUserDevices()

	query := FlumeWaterQuery{
		Bucket:        FlumeWaterBucketDay,
		SinceDatetime: "2021-03-12 00:00:00",
		RequestID:     "test",
	}
	results, err := client.QueryUserDevice(devices.Data[0].ID, FlumeWaterQueryRequest{
		Queries: []FlumeWaterQuery{query},
	})
	assert.Nil(t, err)
	for key := range results.Data[0] {
		assert.Equal(t, "test", key)
	}
}
