package flumewater

import (
	"fmt"
	"log"
)

type FlumeWaterFetchDeviceResponse struct {
	*FlumeResponseBase
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

func (fw *FlumeWaterClient) FetchUserDevices() (flumeResp *FlumeWaterFetchDeviceResponse, err error) {
	if fw.userID == 0 {
		fw.GetToken()
	}

	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/devices"
	flumeResp = new(FlumeWaterFetchDeviceResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}

func (fw *FlumeWaterClient) FetchUserDevice(deviceID string) (flumeResp *FlumeWaterFetchDeviceResponse, err error) {
	if fw.userID == 0 {
		fw.GetToken()
	}

	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/devices/" + deviceID

	flumeResp = new(FlumeWaterFetchDeviceResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}