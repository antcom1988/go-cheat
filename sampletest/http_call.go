package sampletest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var ErrBadStatusCode = errors.New("Bad Status")

type Response struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

func MakeHTTPCall(url string) (*Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrBadStatusCode
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}

	r := &Response{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}

	return r, nil
}
