package flumewater

import (
	"fmt"
	"log"

	"github.com/google/go-querystring/query"
)

type FlumeWaterDeviceBudgetsResponse struct {
	*ResponseBase
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

func (fw *Client) FetchDeviceBudgets(deviceID string, queryParams QueryParamsBase) (budgets []FlumeWaterBudget, err error) {
	if queryParams.SortDirection == "" {
		queryParams.SortDirection = FlumeWaterSortDirectionAsc
	}
	if queryParams.SortField == "" {
		queryParams.SortField = "id"
	}

	v, _ := query.Values(queryParams)
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/devices/" + deviceID + "?" + v.Encode()

	var flumeResp FlumeWaterDeviceBudgetsResponse
	err = fw.FlumeGet(fetchURL, &flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	budgets = flumeResp.Data

	return budgets, nil
}
