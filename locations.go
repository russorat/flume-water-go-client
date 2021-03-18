package flumewater

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

type FlumeWaterLocation struct {
	ID              int                    `json:"id"`
	UserID          int                    `json:"user_id"`
	Name            string                 `json:"name"`
	PrimaryLocation bool                   `json:"primary_location"`
	Address         string                 `json:"address"`
	Address2        string                 `json:"address_2"`
	City            string                 `json:"city"`
	State           string                 `json:"state"`
	PostalCode      string                 `json:"postal_code"`
	Country         string                 `json:"country"`
	TZ              string                 `json:"tz"`
	Installation    string                 `json:"installation"`
	BuildingType    string                 `json:"building_type"`
	UsageProfile    FlumeWaterUsageProfile `json:"usage_profile"`
}

type FlumeWaterLocationsResponse struct {
	*ResponseBase
	Data []FlumeWaterLocation `json:"data"`
}

type LocationsQueryParams struct {
	*QueryParamsBase
	ListShared bool `url:"list_shared,omitempty"`
}

func NewLocationsQueryParams() *LocationsQueryParams {
	return &LocationsQueryParams{
		QueryParamsBase: &QueryParamsBase{},
	}
}

func (fw *Client) FetchUserLocations(queryParams LocationsQueryParams) (locations []FlumeWaterLocation, err error) {
	if queryParams.SortDirection == "" {
		queryParams.SortDirection = FlumeWaterSortDirectionAsc
	}
	if queryParams.SortField == "" {
		queryParams.SortField = "id"
	}

	v, _ := query.Values(queryParams)
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/locations?" + v.Encode()

	var flumeResp FlumeWaterLocationsResponse
	err = fw.FlumeGet(fetchURL, &flumeResp)
	if err != nil {
		log.Fatal(err)
	}
	locations = flumeResp.Data
	return locations, nil
}

func (fw *Client) FetchUserLocation(locationID int) (location FlumeWaterLocation, err error) {
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/locations/" + fmt.Sprint(locationID)

	var flumeResp FlumeWaterLocationsResponse
	err = fw.FlumeGet(fetchURL, &flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	location = flumeResp.Data[0]

	return location, nil
}

type FlumeWaterUpdateUserLocationRequest struct {
	AwayMode bool `json:"away_mode"`
}

func (fw *Client) UpdateUserLocation(locationID string, awayMode bool) (flumeResp *ResponseBase, err error) {
	bodyParams := FlumeWaterUpdateUserLocationRequest{
		AwayMode: awayMode,
	}
	patchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID()) + "/locations/" + fmt.Sprint(locationID)
	jsonValue, _ := json.Marshal(bodyParams)

	req, err := http.NewRequest(http.MethodPatch, patchURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", defaultContentType)

	resp, err := fw.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("when reading from [%s] received status code: %d", patchURL, resp.StatusCode)
	}

	flumeResp = new(ResponseBase)
	decoder := json.NewDecoder(resp.Body)

	if err = decoder.Decode(flumeResp); err != nil {
		return nil, fmt.Errorf("payload JSON decode failed: %w", err)
	}

	return flumeResp, nil
}
