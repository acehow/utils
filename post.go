package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

/*
var client = http.Client{
	Timeout: 10 * time.Second,
}
*/

func HttpPostJson(client *http.Client, url string, data interface{}, result interface{}, header map[string]string) error {
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	if header != nil {
		for k, v := range header {
			request.Header.Add(k, v)
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&result); err != nil {
		return err
	}

	return nil
}

func HttpPostForm(client *http.Client, posturl string, data url.Values, result interface{}, host string) error {
	request, err := http.NewRequest(http.MethodPost, posturl, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//request.Host = host

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&result); err != nil {
		return err
	}

	return nil
}

func HttpGetRequest(client *http.Client, url string, result interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)

	return err
}
