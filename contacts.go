package flumewater

import (
	"log"

	"github.com/google/go-querystring/query"
)

type FlumeWaterContactParams struct {
	*QueryParamsBase
	Type     FlumeWaterContactType     `json:"type,omitempty"`
	Category FlumeWaterContactCategory `json:"category,omitempty"`
}

func NewFlumeWaterContactParams() *FlumeWaterContactParams {
	return &FlumeWaterContactParams{
		QueryParamsBase: &QueryParamsBase{},
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
	*ResponseBase
	Data []FlumeWaterContact `json:"data"`
}

type FlumeWaterContact struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Detail   string `json:"detail"`
}

func (fw *Client) FetchFlumeContactInfo(queryParams FlumeWaterContactParams) (contacts []FlumeWaterContact, err error) {
	v, _ := query.Values(queryParams)
	var encodedValues = v.Encode()
	fetchURL := baseURL + "/contacts"
	if len(encodedValues) > 0 {
		fetchURL += "?" + v.Encode()
	}

	var flumeResp FlumeWaterContactResponse
	err = fw.FlumeGet(fetchURL, &flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	contacts = flumeResp.Data

	return contacts, nil
}
