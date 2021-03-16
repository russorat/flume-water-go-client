package flumewater

import (
	"fmt"
	"log"
)

type FlumeWaterFetchUserResponse struct {
	*FlumeResponseBase
	Data []FlumeWaterUser `json:"data"`
}

type FlumeWaterUser struct {
	ID             int    `json:"id"`
	EmailAddress   string `json:"email_address"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Phone          string `json:"phone"`
	Status         string `json:"status"`
	Type           string `json:"type"`
	SignupDatetime string `json:"signup_datetime,omitempty"`
}

func (fw *FlumeWaterClient) FetchUser() (flumeResp *FlumeWaterFetchUserResponse, err error) {
	if fw.userID == 0 {
		fw.GetToken()
	}

	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.userID)

	flumeResp = new(FlumeWaterFetchUserResponse)
	err = fw.FlumeGet(fetchURL, flumeResp)
	if err != nil {
		log.Fatal(err)
	}

	return flumeResp, nil
}
