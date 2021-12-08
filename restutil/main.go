package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
)

type HttpBinGetResp struct {
	Args struct {
	} `json:"args"`
	Headers struct {
		Accept         string `json:"Accept"`
		AcceptEncoding string `json:"Accept-Encoding"`
		AcceptLanguage string `json:"Accept-Language"`
		Dnt            string `json:"Dnt"`
		Host           string `json:"Host"`
		Referer        string `json:"Referer"`
		SecFetchDest   string `json:"Sec-Fetch-Dest"`
		SecFetchMode   string `json:"Sec-Fetch-Mode"`
		SecFetchSite   string `json:"Sec-Fetch-Site"`
		SecGpc         string `json:"Sec-Gpc"`
		UserAgent      string `json:"User-Agent"`
		XAmznTraceID   string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func main() {

	var httpBinGetResp HttpBinGetResp
	reqURL, err := url.Parse("https://httpbin.org/get")
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse URL")
	}

	t := &http.Transport{}

	_, err = restUtil(reqURL.String(), "GET", t, nil, &httpBinGetResp)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get response")
		return
	}
	log.Info().Msgf("%+v", httpBinGetResp)

}
func restUtil(reqURL, method string, t *http.Transport, reqBody io.Reader, respValue interface{}) (int, error) {
	client := &http.Client{Transport: t}

	req, err := http.NewRequest(method, reqURL, nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusAccepted:
	case http.StatusCreated:
	default:
		return resp.StatusCode, err
	}

	err = json.NewDecoder(resp.Body).Decode(respValue)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return resp.StatusCode, nil
}
