package flumewater

import (
	"fmt"
	"log"
	"time"

	"github.com/google/go-querystring/query"
)

type FlumeWaterUserSubscriptionsResponse struct {
	*ResponseBase
	Data []FlumeWaterSubscription `json:"data"`
}

type FlumeWaterSubscription struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	AlertType         string    `json:"alert_type"`
	AlertInfo         string    `json:"alert_info"`
	DeviceID          string    `json:"device_id"`
	NotificationTypes int       `json:"notification_types"`
	CreatedDatetime   time.Time `json:"created_datetime"`
	UpdatedDatetime   time.Time `json:"updated_datetime"`
}

type FlumeWaterSubscriptionParams struct {
	*QueryParamsBase
	AlertType         string                     `json:"alert_type,omitempty"`
	NotificationTypes int                        `json:"notification_types,omitempty"`
	NotificationType  FlumeWaterNotificationType `json:"notification_type,omitempty"`
	DeviceID          string                     `json:"device_id,omitempty"`
	DeviceType        FlumeWaterDeviceType       `json:"device_type,omitempty"`
	LocationID        int                        `json:"location_id,omitempty"`
}

func NewFlumeWaterSubscriptionParams() *FlumeWaterSubscriptionParams {
	return &FlumeWaterSubscriptionParams{
		QueryParamsBase: &QueryParamsBase{},
	}
}

func (fw *Client) FetchUserSubscriptions(queryParams FlumeWaterSubscriptionParams) (subscriptions []FlumeWaterSubscription, err error) {
	if queryParams.SortDirection == "" {
		queryParams.SortDirection = FlumeWaterSortDirectionAsc
	}
	if queryParams.SortField == "" {
		queryParams.SortField = "id"
	}

	v, _ := query.Values(queryParams)
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/subscriptions?" + v.Encode()

	var flumeResp FlumeWaterUserSubscriptionsResponse
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	subscriptions = flumeResp.Data

	return subscriptions, nil
}

func (fw *Client) FetchUserSubscription(subscriptionID int) (subscription FlumeWaterSubscription, err error) {
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/subscriptions/" + fmt.Sprint(subscriptionID)

	var flumeResp FlumeWaterUserSubscriptionsResponse
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	subscription = flumeResp.Data[0]

	return subscription, nil
}
