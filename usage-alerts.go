package flumewater

import (
	"fmt"
	"log"
	"time"

	"github.com/google/go-querystring/query"
)

type FlumeWaterUsageAlertsResponse struct {
	*ResponseBase
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
	*QueryParamsBase
	DeviceID  string `url:"device_id,omitempty"`
	FlumeLeak bool   `url:"flume_leak,omitempty"`
}

func NewFlumeWaterUsageAlertsParams() *FlumeWaterUsageAlertsParams {
	return &FlumeWaterUsageAlertsParams{
		QueryParamsBase: &QueryParamsBase{},
	}
}

func (fw *Client) FetchUserUsageAlerts(queryParams FlumeWaterUsageAlertsParams) (flumeResp *FlumeWaterUsageAlertsResponse, err error) {
	if queryParams.SortDirection == "" {
		queryParams.SortDirection = FlumeWaterSortDirectionAsc
	}
	if queryParams.SortField == "" {
		queryParams.SortField = "id"
	}

	v, _ := query.Values(queryParams)
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/usage-alerts?" + v.Encode()

	flumeResp = new(FlumeWaterUsageAlertsResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}

type FlumeWaterUsageAlertRuleResponse struct {
	*ResponseBase
	Data []FlumeWaterUsageAlertRule `json:"data"`
}

type FlumeWaterUsageAlertRule struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Active      bool    `json:"active"`
	FlowRate    float64 `json:"flow_rate"`
	Duration    int     `json:"duration"`
	NotifyEvery int     `json:"notify_every"`
}

func (fw *Client) FetchAllUsageAlertRulesForDevice(deviceID string, queryParams QueryParamsBase) (alertRules []FlumeWaterUsageAlertRule, err error) {
	if queryParams.SortDirection == "" {
		queryParams.SortDirection = FlumeWaterSortDirectionAsc
	}
	if queryParams.SortField == "" {
		queryParams.SortField = "id"
	}

	v, _ := query.Values(queryParams)
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/devices/" + deviceID + "/rules/usage-alerts?" + v.Encode()

	var flumeResp FlumeWaterUsageAlertRuleResponse
	err = fw.FlumeGet(fetchURL, &flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	alertRules = flumeResp.Data

	return alertRules, nil
}

func (fw *Client) FetchSingleUsageAlertRulesForDevice(deviceID string, RuleID int) (usageAlert FlumeWaterUsageAlertRule, err error) {
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/devices/" + deviceID + "/rules/usage-alerts/" + fmt.Sprint(RuleID)

	var flumeResp FlumeWaterUsageAlertRuleResponse
	err = fw.FlumeGet(fetchURL, &flumeResp)
	if err != nil {
		log.Fatal(err)
	}
	usageAlert = flumeResp.Data[0]

	return usageAlert, nil
}
