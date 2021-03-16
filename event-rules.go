package flumewater

import (
	"fmt"
	"log"

	"github.com/google/go-querystring/query"
)

type FlumeWaterEventRuleResponse struct {
	*FlumeResponseBase
	Data []FlumeWaterEventRule `json:"data"`
}

type FlumeWaterEventRule struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Active           bool    `json:"active"`
	FlowRate         float64 `json:"flow_rate"`
	Duration         int     `json:"duration"`
	NotifyEvery      int     `json:"notify_every"`
	NotificationType string  `json:"notification_type"`
}

func (fw *FlumeWaterClient) FetchAllEventRulesForDevice(deviceID string, queryParams BaseQueryParams) (flumeResp *FlumeWaterEventRuleResponse, err error) {
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
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/devices/" + deviceID + "/rules?" + v.Encode()

	flumeResp = new(FlumeWaterEventRuleResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}
