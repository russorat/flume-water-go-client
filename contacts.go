package flumewater

import (
	"log"

	"github.com/google/go-querystring/query"
)

type FlumeWaterContactParams struct {
	*BaseQueryParams
	Type     FlumeWaterContactType     `json:"type,omitempty"`
	Category FlumeWaterContactCategory `json:"category,omitempty"`
}

func NewFlumeWaterContactParams() *FlumeWaterContactParams {
	return &FlumeWaterContactParams{
		BaseQueryParams: &BaseQueryParams{},
	}
}

type FlumeWaterContactType string

const (
	FlumeWaterContactTypePhone FlumeWaterContactType = "PHONE"
	FlumeWaterContactTypeEmail FlumeWaterContactType = "EMAIL"
)

type FlumeWaterContactCategory string

const (
	FlumeWaterContactCategorySupport FlumeWaterContactCategory = "SUPPORT"
)

type FlumeWaterContactResponse struct {
	*FlumeResponseBase
	Data []FlumeWaterContact `json:"data"`
}

type FlumeWaterContact struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Detail   string `json:"detail"`
}

func (fw *FlumeWaterClient) FetchFlumeContactInfo(queryParams FlumeWaterContactParams) (flumeResp *FlumeWaterContactResponse, err error) {
	v, _ := query.Values(queryParams)
	var encodedValues = v.Encode()
	fetchURL := baseURL + "/contacts"
	if len(encodedValues) > 0 {
		fetchURL += "?" + v.Encode()
	}

	flumeResp = new(FlumeWaterContactResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}
