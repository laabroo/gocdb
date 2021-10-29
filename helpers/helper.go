package helpers

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string, params map[string]string) (interface{}, error) {
	var client = &http.Client{}
	var err error

	targetURL := url
	if params != nil {
		index := 0
		for key, value := range params {
			if index == 0 {
				targetURL += "?" + key + "=" + value
			} else {
				targetURL += "&" + key + "=" + value
			}
			index += 1
		}
	}

	request, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return string(data), nil
}
