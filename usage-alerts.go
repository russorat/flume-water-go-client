package flumewater

import (
	"fmt"
	"log"
	"time"

	"github.com/google/go-querystring/query"
)

type FlumeWaterUsageAlertsResponse struct {
	*FlumeResponseBase
	Data []FlumeWaterUsageAlert `json:"data"`
}

type FlumeWaterUsageAlert struct {
	ID                int                       `json:"id"`
	DeviceID          string                    `json:"device_id"`
	TriggeredDatetime time.Time                 `json:"triggered_datetime"`
	FlumeLeak         bool                      `json:"flume_leak"`
	Query             FlumeWaterUsageAlertQuery `json:"query"`
	EventRuleName     string                    `json:"event_rule_name"`
}

type FlumeWaterUsageAlertQuery struct {
	RequestID     string   `json:"request_id"`
	SinceDatetime string   `json:"since_datetime"`
	UntilDatetime string   `json:"until_datetime"`
	Tz            string   `json:"tz"`
	Bucket        string   `json:"bucket"`
	DeviceID      []string `json:"device_id"`
}

type FlumeWaterUsageAlertsParams struct {
	*BaseQueryParams
	DeviceID  string `json:"device_id,omitempty"`
	FlumeLeak bool   `json:"flume_leak,omitempty"`
}

func NewFlumeWaterUsageAlertsParams() *FlumeWaterUsageAlertsParams {
	return &FlumeWaterUsageAlertsParams{
		BaseQueryParams: &BaseQueryParams{},
	}
}

func (fw *FlumeWaterClient) FetchUserUsageAlerts(queryParams FlumeWaterUsageAlertsParams) (flumeResp *FlumeWaterUsageAlertsResponse, err error) {
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
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/usage-alerts?" + v.Encode()

	flumeResp = new(FlumeWaterUsageAlertsResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}

type FlumeWaterUsageAlertRuleResponse struct {
	*FlumeResponseBase
	Data []FlumeWaterEventRule `json:"data"`
}

type FlumeWaterUsageAlertRule struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Active      bool    `json:"active"`
	FlowRate    float64 `json:"flow_rate"`
	Duration    int     `json:"duration"`
	NotifyEvery int     `json:"notify_every"`
}

func (fw *FlumeWaterClient) FetchAllUsageAlertRulesForDevice(deviceID string, queryParams BaseQueryParams) (flumeResp *FlumeWaterUsageAlertRuleResponse, err error) {
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
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/devices/" + deviceID + "/rules/usage-alerts?" + v.Encode()

	flumeResp = new(FlumeWaterUsageAlertRuleResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}

func (fw *FlumeWaterClient) FetchSingleUsageAlertRulesForDevice(deviceID string, RuleID int) (flumeResp *FlumeWaterUsageAlertRuleResponse, err error) {
	if fw.userID == 0 {
		fw.GetToken()
	}

	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/devices/" + deviceID + "/rules/usage-alerts/" + fmt.Sprint(RuleID)

	flumeResp = new(FlumeWaterUsageAlertRuleResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}
