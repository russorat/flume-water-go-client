package flumewater

import (
	"fmt"
	"log"
	"time"

	"github.com/google/go-querystring/query"
)

type FlumeWaterNotificationType int

const (
	FlumeWaterNotificationTypeUsageAlert FlumeWaterNotificationType = 1
	FlumeWaterNotificationTypeBudget     FlumeWaterNotificationType = 2
	FlumeWaterNotificationTypeGeneral    FlumeWaterNotificationType = 4
	FlumeWaterNotificationTypeHeartbeat  FlumeWaterNotificationType = 8
	FlumeWaterNotificationTypeBattery    FlumeWaterNotificationType = 16
)

type FlumeWaterNotificationsResponse struct {
	*ResponseBase
	Data []FlumeWaterNotification `json:"data"`
}

type FlumeWaterNotification struct {
	ID              int                          `json:"id"`
	DeviceID        string                       `json:"device_id"`
	UserID          int                          `json:"user_id"`
	EventRule       string                       `json:"event_rule"`
	Type            int                          `json:"type"`
	Message         string                       `json:"message"`
	CreatedDatetime time.Time                    `json:"created_datetime"`
	Title           string                       `json:"title"`
	Read            bool                         `json:"read"`
	Extra           FlumeWaterNotificationsExtra `json:"extra"`
}

type FlumeWaterNotificationsExtra struct {
	Percentage    int       `json:"percentage"`
	BudgetType    string    `json:"budget_type"`
	BudgetStart   time.Time `json:"budget_start"`
	EventRuleName string    `json:"event_rule_name"`
}

type FlumeWaterNotificationsParams struct {
	*QueryParamsBase
	DeviceID          string                     `url:"device_id,omitempty"`
	LocationID        int                        `url:"location_id,omitempty"`
	AlertType         string                     `url:"alert_type,omitempty"`
	NotificationType  FlumeWaterNotificationType `url:"type,omitempty"`
	NotificationTypes int                        `url:"types,omitempty"`
	Read              bool                       `url:"read,omitempty"`
}

func NewFlumeWaterNotificationsParams() *FlumeWaterNotificationsParams {
	return &FlumeWaterNotificationsParams{
		QueryParamsBase: &QueryParamsBase{},
	}
}

func (fw *Client) FetchUserNotifications(queryParams FlumeWaterNotificationsParams) (notifications []FlumeWaterNotification, err error) {
	if queryParams.SortDirection == "" {
		queryParams.SortDirection = FlumeWaterSortDirectionAsc
	}
	if queryParams.SortField == "" {
		queryParams.SortField = "id"
	}

	v, _ := query.Values(queryParams)
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/notifications?" + v.Encode()

	var flumeResp FlumeWaterNotificationsResponse
	err = fw.FlumeGet(fetchURL, &flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	notifications = flumeResp.Data

	return notifications, nil
}
