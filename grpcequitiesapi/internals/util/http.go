package util

import (
	"io/ioutil"
	"net/http"
)

func GetHttp(url string) ([]byte, int, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, 500, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, 500, err
	}

	return body, res.StatusCode, nil
}
