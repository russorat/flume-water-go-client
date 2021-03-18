package flumewater

import (
	"fmt"
	"log"

	"github.com/google/go-querystring/query"
)

type FlumeWaterFetchDeviceResponse struct {
	*ResponseBase
	Data []FlumeWaterDevice `json:"data"`
}

type FlumeWaterDevice struct {
	ID            string               `json:"id"`
	BridgeID      string               `json:"bridge_id"`
	Type          FlumeWaterDeviceType `json:"type"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	AddedDatetime string               `json:"added_datetime"`
	UserID        int                  `json:"user_id"`
	Registered    bool                 `json:"registered"`
	Oriented      bool                 `json:"oriented"`
	LastSeen      string               `json:"last_seen"`
	Location      FlumeWaterLocation   `json:"location"`
	User          FlumeWaterUser       `json:"user"`
}

type FlumeWaterDeviceType int

const (
	FlumeWaterDeviceTypeBridge FlumeWaterDeviceType = 1
	FlumeWaterDeviceTypeSensor FlumeWaterDeviceType = 2
)

type FlumeWaterUsageProfile struct {
	ID                 int    `json:"id"`
	Score              int    `json:"score"`
	Residents          string `json:"residents"`
	Bathrooms          string `json:"bathrooms"`
	Irrigation         string `json:"irrigation"`
	IrrigationFreq     string `json:"irrigation_freq"`
	IrrigationMaxCycle int    `json:"irrigation_max_cycle"`
}

type FlumeWaterFetchDeviceRequest struct {
	IncludeUser     bool `json:"user,omitempty"`
	IncludeLocation bool `json:"location,omitempty"`
}

func (fw *Client) FetchUserDevices(queryParams FlumeWaterFetchDeviceRequest) (devices []FlumeWaterDevice, err error) {
	v, _ := query.Values(queryParams)
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/devices?" + v.Encode()

	var flumeResp FlumeWaterFetchDeviceResponse
	err = fw.FlumeGet(fetchURL, &flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	devices = flumeResp.Data

	return devices, nil
}

func (fw *Client) FetchUserDevice(deviceID string, queryParams FlumeWaterFetchDeviceRequest) (device FlumeWaterDevice, err error) {
	v, _ := query.Values(queryParams)
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/devices/" + deviceID + "?" + v.Encode()

	var flumeResp FlumeWaterFetchDeviceResponse
	err = fw.FlumeGet(fetchURL, &flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	device = flumeResp.Data[0]

	return device, nil
}
