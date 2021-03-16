package flumewater

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type FlumeWaterAuthRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type FlumeWaterAuthResponse struct {
	*FlumeResponseBase
	Data []FlumeWaterAuthData `json:"data"`
}

type FlumeWaterAuthData struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func (fw *FlumeWaterClient) GetToken() (err error) {
	authParams := FlumeWaterAuthRequest{
		GrantType:    "password",
		ClientID:     fw.ClientID,
		ClientSecret: fw.ClientSecret,
		Username:     fw.Username,
		Password:     fw.Password,
	}
	tokenURL := baseURL + "/oauth/token"

	jsonValue, _ := json.Marshal(authParams)

	req, err := http.NewRequest(http.MethodPost, tokenURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", defaultContentType)

	resp, err := fw.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("when reading from [%s] received status code: %d", tokenURL, resp.StatusCode)
	}

	var flumeResp = new(FlumeWaterAuthResponse)
	decoder := json.NewDecoder(resp.Body)

	if err = decoder.Decode(flumeResp); err != nil {
		return fmt.Errorf("payload JSON decode failed: %w", err)
	}
	fw.AuthData = flumeResp.Data[0]

	type MyCustomClaims struct {
		UserID int `json:"user_id"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(fw.AuthData.AccessToken, &MyCustomClaims{}, nil)
	if token == nil {
		return err
	}
	claims, _ := token.Claims.(*MyCustomClaims)
	fw.userID = claims.UserID
	return nil
}

type FlumeWaterRefreshTokenRequest struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (fw *FlumeWaterClient) RefreshToken() (err error) {
	refreshParams := FlumeWaterRefreshTokenRequest{
		GrantType:    "refresh_token",
		RefreshToken: fw.AuthData.RefreshToken,
		ClientID:     fw.ClientID,
		ClientSecret: fw.ClientSecret,
	}
	tokenURL := baseURL + "/oauth/token"
	jsonValue, _ := json.Marshal(refreshParams)

	req, err := http.NewRequest(http.MethodPost, tokenURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", defaultContentType)

	resp, err := fw.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("when reading from [%s] received status code: %d", tokenURL, resp.StatusCode)
	}

	var flumeResp = new(FlumeWaterAuthResponse)
	decoder := json.NewDecoder(resp.Body)

	if err = decoder.Decode(flumeResp); err != nil {
		return fmt.Errorf("payload JSON decode failed: %w", err)
	}
	fw.AuthData = flumeResp.Data[0]

	return nil
}
