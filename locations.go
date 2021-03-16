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
	*FlumeResponseBase
	Data []FlumeWaterLocation `json:"data"`
}

type LocationsQueryParams struct {
	*BaseQueryParams
	ListShared bool `url:"list_shared,omitempty"`
}

func NewLocationsQueryParams() *LocationsQueryParams {
	return &LocationsQueryParams{
		BaseQueryParams: &BaseQueryParams{},
	}
}

func (fw *FlumeWaterClient) FetchUserLocations(queryParams LocationsQueryParams) (flumeResp *FlumeWaterLocationsResponse, err error) {
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
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/locations?" + v.Encode()

	flumeResp = new(FlumeWaterLocationsResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}

func (fw *FlumeWaterClient) FetchUserLocation(locationID int) (flumeResp *FlumeWaterLocationsResponse, err error) {
	if fw.userID == 0 {
		fw.GetToken()
	}

	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/locations/" + fmt.Sprint(locationID)

	flumeResp = new(FlumeWaterLocationsResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}

type FlumeWaterUpdateUserLocationRequest struct {
	AwayMode bool `json:"away_mode"`
}

func (fw *FlumeWaterClient) UpdateUserLocation(locationID string, awayMode bool) (flumeResp *FlumeResponseBase, err error) {
	bodyParams := FlumeWaterUpdateUserLocationRequest{
		AwayMode: awayMode,
	}
	patchURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/locations/" + fmt.Sprint(locationID)
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

	flumeResp = new(FlumeResponseBase)
	decoder := json.NewDecoder(resp.Body)

	if err = decoder.Decode(flumeResp); err != nil {
		return nil, fmt.Errorf("payload JSON decode failed: %w", err)
	}

	return flumeResp, nil
}
