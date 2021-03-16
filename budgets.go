package flumewater

import (
	"fmt"
	"log"

	"github.com/google/go-querystring/query"
)

type FlumeWaterDeviceBudgetsResponse struct {
	*FlumeResponseBase
	Data []FlumeWaterBudget `json:"data"`
}

type FlumeWaterBudget struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Value      int    `json:"value"`
	Thresholds []int  `json:"thresholds"`
	Actual     int    `json:"actual"`
}

func (fw *FlumeWaterClient) FetchDeviceBudgets(deviceID string, queryParams BaseQueryParams) (flumeResp *FlumeWaterDeviceBudgetsResponse, err error) {
	if fw.userID == 0 {
		fw.GetToken()
	}

	if queryParams.SortDirection == "" {
		queryParams.SortDirection = FlumeWaterSortDirectionAsc
	}
	if queryParams.SortField == "" {
		queryParams.SortField = "id"
	}

	v, _ := query.Values(queryParams)
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/devices/" + deviceID + "?" + v.Encode()

	flumeResp = new(FlumeWaterDeviceBudgetsResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}
