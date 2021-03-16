package flumewater

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	defaultClientTimeout = 5 * time.Second
	defaultContentType   = "application/json; charset=utf-8"
	baseURL              = "https://api.flumewater.com"
	userAgent            = "github.com/russorat/telegraf-flume-water-input"
)

type FlumeWaterClient struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
	Timeout      time.Duration

	AuthData FlumeWaterAuthData
	userID   int
	client   *http.Client
}

func NewClient(clientID string, clientSecret string, username string, password string) (client FlumeWaterClient) {
	client = FlumeWaterClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
		Timeout:      defaultClientTimeout,
		client:       http.DefaultClient,
	}

	return client
}

type FlumeResponseBase struct {
	Success     bool     `json:"success"`
	Code        int      `json:"code"`
	Message     string   `json:"message"`
	HTTPCode    int      `json:"http_code"`
	HTTPMessage string   `json:"http_message"`
	Detailed    []string `json:"detailed"`
	Count       int      `json:"count"`
	Pagination  struct {
		Next string `json:"next"`
		Prev string `json:"prev"`
	} `json:"pagination"`
}

type BaseQueryParams struct {
	Limit         int                     `url:"limit,omitempty"`
	Offset        int                     `url:"offset,omitempty"`
	SortField     string                  `url:"sort_field,omitempty"`
	SortDirection FlumeWaterSortDirection `url:"sort_direction,omitempty"`
}

type FlumeWaterErrorResponse struct {
	*FlumeResponseBase
}

type FlumeWaterDetailed struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (fw *FlumeWaterClient) Connect() error {
	if fw.Timeout == 0 {
		fw.Timeout = defaultClientTimeout
	}

	ctx := context.Background()
	client, err := fw.createClient(ctx)
	if err != nil {
		return err
	}

	fw.client = client

	return nil
}

func (fw *FlumeWaterClient) createClient(ctx context.Context) (*http.Client, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
		Timeout: fw.Timeout,
	}

	return client, nil
}

func (fw *FlumeWaterClient) FlumeGet(fetchURL string, flumeResponse interface{}) (err error) {
	if fw.AuthData.AccessToken == "" {
		fw.GetToken()
	}

	req, err := http.NewRequest(http.MethodGet, fetchURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", defaultContentType)
	req.Header.Set("Authorization", "Bearer "+fw.AuthData.AccessToken)

	resp, err := fw.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)

		var flumeErrorResp = new(FlumeWaterErrorResponse)
		decoder := json.NewDecoder(resp.Body)

		if err = decoder.Decode(flumeErrorResp); err != nil {
			return fmt.Errorf("payload JSON decode failed: %w", err)
		}
		return fmt.Errorf("when reading from [%s] received status code: %d with error: %s", fetchURL, resp.StatusCode, bodyString)
	}

	decoder := json.NewDecoder(resp.Body)

	if err = decoder.Decode(flumeResponse); err != nil {
		return fmt.Errorf("payload JSON decode failed: %w", err)
	}

	return nil
}
