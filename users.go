package flumewater

import (
	"fmt"
	"log"
)

type FlumeWaterFetchUserResponse struct {
	*ResponseBase
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

func (fw *Client) FetchUser() (flumeUser FlumeWaterUser, err error) {
	fetchURL := baseURL + "/users/" + fmt.Sprint(fw.UserID())

	var flumeResp FlumeWaterFetchUserResponse
	err = fw.FlumeGet(fetchURL, &flumeResp)
	if err != nil {
		log.Fatal(err)
	}
	flumeUser = flumeResp.Data[0]
	return flumeUser, nil
}
