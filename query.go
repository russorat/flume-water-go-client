package flumewater

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type FlumeWaterQueryResponse struct {
	*FlumeResponseBase
	Data []FlumeWaterQueryResult `json:"data"`
}

type FlumeWaterQueryResultBucket struct {
	Datetime string  `json:"datetime"`
	Value    float64 `json:"value"`
}

type FlumeWaterQueryResult map[string][]FlumeWaterQueryResultBucket

type FlumeWaterQueryRequest struct {
	Queries []FlumeWaterQuery `json:"queries"`
}

type FlumeWaterQuery struct {
	Bucket          FlumeWaterBucket        `json:"bucket"`
	GroupMultiplier int                     `json:"group_multiplier,omitempty"`
	SinceDatetime   string                  `json:"since_datetime"`
	UntilDatetime   string                  `json:"until_datetime,omitempty"`
	Operation       FlumeWaterOperation     `json:"operation,omitempty"`
	Units           FlumeWaterUnit          `json:"units,omitempty"`
	SortDirection   FlumeWaterSortDirection `json:"sort_direction,omitempty"`
	RequestID       string                  `json:"request_id"`
}

type FlumeWaterUnit string

const (
	FlumeWaterUnitGallon      FlumeWaterUnit = "GALLONS"
	FlumeWaterUnitLiters      FlumeWaterUnit = "LITERS"
	FlumeWaterUnitCubicFeet   FlumeWaterUnit = "CUBIC_FEET"
	FlumeWaterUnitCubicMeters FlumeWaterUnit = "CUBIC_METERS"
)

type FlumeWaterBucket string

const (
	FlumeWaterBucketMinute FlumeWaterBucket = "MIN"
	FlumeWaterBucketHour   FlumeWaterBucket = "HR"
	FlumeWaterBucketDay    FlumeWaterBucket = "DAY"
	FlumeWaterBucketMonth  FlumeWaterBucket = "MON"
	FlumeWaterBucketYear   FlumeWaterBucket = "YR"
)

type FlumeWaterOperation string

const (
	FlumeWaterBucketSum     FlumeWaterOperation = "SUM"
	FlumeWaterBucketAverage FlumeWaterOperation = "AVG"
	FlumeWaterBucketMin     FlumeWaterOperation = "MIN"
	FlumeWaterBucketMax     FlumeWaterOperation = "MAX"
	FlumeWaterBucketCount   FlumeWaterOperation = "CNT"
)

type FlumeWaterSortDirection string

const (
	FlumeWaterSortDirectionAsc  FlumeWaterSortDirection = "ASC"
	FlumeWaterSortDirectionDesc FlumeWaterSortDirection = "DESC"
)

type FlumeWaterQueryErrorResponse struct {
	*FlumeResponseBase
	Detailed []FlumeWaterDetailed `json:"detailed"`
}

func (fw *FlumeWaterClient) QueryUserDevice(deviceID string, Queries FlumeWaterQueryRequest) (flumeResp *FlumeWaterQueryResponse, err error) {
	queryDeviceURL := baseURL + "/users/" + fmt.Sprint(fw.userID) + "/devices/" + deviceID + "/query"
	jsonValue, _ := json.Marshal(Queries)

	req, err := http.NewRequest(http.MethodPost, queryDeviceURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", defaultContentType)
	if fw.AuthData.AccessToken == "" {
		fw.GetToken()
	}
	req.Header.Set("Authorization", "Bearer "+fw.AuthData.AccessToken)

	resp, err := fw.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var flumeErrorResp = new(FlumeWaterErrorResponse)
		decoder := json.NewDecoder(resp.Body)

		if err = decoder.Decode(flumeErrorResp); err != nil {
			return nil, fmt.Errorf("payload JSON decode failed: %w", err)
		}
		return nil, fmt.Errorf("when reading from [%s] received status code: %d with error: %s", queryDeviceURL, resp.StatusCode, strings.Join(flumeErrorResp.Detailed, ","))
	}

	flumeResp = new(FlumeWaterQueryResponse)
	decoder := json.NewDecoder(resp.Body)

	if err = decoder.Decode(flumeResp); err != nil {
		return nil, fmt.Errorf("payload JSON decode failed: %w", err)
	}

	return flumeResp, nil
}
