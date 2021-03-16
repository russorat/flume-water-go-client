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
	*BaseQueryParams
	DeviceID          string                     `json:"device_id,omitempty"`
	LocationID        int                        `json:"location_id,omitempty"`
	AlertType         string                     `json:"alert_type,omitempty"`
	NotificationType  FlumeWaterNotificationType `json:"type,omitempty"`
	NotificationTypes int                        `json:"types,omitempty"`
	Read              bool                       `json:"read,omitempty"`
}

func NewFlumeWaterNotificationsParams() *FlumeWaterNotificationsParams {
	return &FlumeWaterNotificationsParams{
		BaseQueryParams: &BaseQueryParams{},
	}
}

func (fw *FlumeWaterClient) FetchUserNotifications(queryParams FlumeWaterNotificationsParams) (flumeResp *FlumeWaterNotificationsResponse, err error) {
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
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/notifications?" + v.Encode()

	flumeResp = new(FlumeWaterNotificationsResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}