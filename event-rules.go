package flumewater

import (
	"fmt"
	"log"

	"github.com/google/go-querystring/query"
)

type FlumeWaterEventRuleResponse struct {
	*ResponseBase
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

func (fw *Client) FetchAllEventRulesForDevice(deviceID string, queryParams QueryParamsBase) (eventRules []FlumeWaterEventRule, err error) {
	if queryParams.SortDirection == "" {
		queryParams.SortDirection = FlumeWaterSortDirectionAsc
	}
	if queryParams.SortField == "" {
		queryParams.SortField = "id"
	}

	v, _ := query.Values(queryParams)
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/devices/" + deviceID + "/rules?" + v.Encode()

	var flumeResp FlumeWaterEventRuleResponse
	err = fw.FlumeGet(fetchURL, &flumeResp)
	if err != nil {
		log.Fatal(err)
	}
	eventRules = flumeResp.Data

	return eventRules, nil
}
