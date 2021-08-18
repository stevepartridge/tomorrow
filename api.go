package tomorrow

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	baseAPIv4URL = "https://api.tomorrow.io/v4"

	HeaderRateLimitDay      = "X-RateLimit-Limit-day"
	HeaderRateLimitHour     = "X-RateLimit-Limit-hour"
	HeaderRateRemainingDay  = "X-RateLimit-Remaining-day"
	HeaderRateRemainingHour = "X-RateLimit-Remaining-hour"
)

func (c *Client) call(method, path string, payload interface{}, query map[string]string) (Response, error) {

	if strings.Index(path, "/") != 0 {
		return Response{}, errors.New("path must start with a forward slash: ' / ' ")
	}

	rawURL := baseAPIv4URL + path + "?apikey=" + c.apiKey

	for k, v := range query {
		rawURL = fmt.Sprintf("%s&%s=%s", rawURL, k, v)
	}

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return Response{}, nil
	}

	for k, v := range query {
		if u.Query().Get(k) != "" {
			u.Query().Set(k, v)
			continue
		}
		u.Query().Add(k, v)
	}

	req := http.Request{
		Method: method,
		URL:    u,
		Header: http.Header{},
	}

	if payload != nil {
		body, err := json.Marshal(payload)
		if err != nil {
			return Response{}, err
		}

		req.Body = ioutil.NopCloser(bytes.NewReader(body))

		req.Header.Add("Content-Type", "application/json")
	}

	fmt.Println("call:", req.URL.String())

	resp, err := client.Do(&req)
	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}
	result := Response{}

	c.RateLimitDay, _ = strconv.Atoi(resp.Header.Get(HeaderRateLimitDay))
	c.RateLimitHour, _ = strconv.Atoi(resp.Header.Get(HeaderRateLimitHour))
	c.RateRemainingDay, _ = strconv.Atoi(resp.Header.Get(HeaderRateRemainingDay))
	c.RateRemainingHour, _ = strconv.Atoi(resp.Header.Get(HeaderRateRemainingHour))

	fmt.Printf("%+v\n", result)

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, err
}
