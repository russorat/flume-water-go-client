package flumewater

import (
	"fmt"
	"log"
	"time"

	"github.com/google/go-querystring/query"
)

type FlumeWaterUserSubscriptionsResponse struct {
	*FlumeResponseBase
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
	*BaseQueryParams
	AlertType         string                     `json:"alert_type,omitempty"`
	NotificationTypes int                        `json:"notification_types,omitempty"`
	NotificationType  FlumeWaterNotificationType `json:"notification_type,omitempty"`
	DeviceID          string                     `json:"device_id,omitempty"`
	DeviceType        FlumeWaterDeviceType       `json:"device_type,omitempty"`
	LocationID        int                        `json:"location_id,omitempty"`
}

func NewFlumeWaterSubscriptionParams() *FlumeWaterSubscriptionParams {
	return &FlumeWaterSubscriptionParams{
		BaseQueryParams: &BaseQueryParams{},
	}
}

func (fw *FlumeWaterClient) FetchUserSubscriptions(queryParams FlumeWaterSubscriptionParams) (flumeResp *FlumeWaterUserSubscriptionsResponse, err error) {
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
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/subscriptions?" + v.Encode()

	flumeResp = new(FlumeWaterUserSubscriptionsResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}

func (fw *FlumeWaterClient) FetchUserSubscription(subscriptionID int) (flumeResp *FlumeWaterUserSubscriptionsResponse, err error) {
	if fw.userID == 0 {
		fw.GetToken()
	}

	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/subscriptions/" + fmt.Sprint(subscriptionID)

	flumeResp = new(FlumeWaterUserSubscriptionsResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}
