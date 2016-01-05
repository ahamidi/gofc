package gofc

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	PERSON_API_ENDPOINT = "https://api.fullcontact.com/v2/person.json"
)

type FCClient struct {
	apiKey string
}

type PersonAPI struct {
	apiKey string
	url    string
}

func NewClient(apiKey string) *FCClient {
	return &FCClient{apiKey}
}

func (fc *FCClient) Person() *PersonAPI {
	return &PersonAPI{fc.apiKey, PERSON_API_ENDPOINT}
}

func getPersonBy(p *PersonAPI, method string, value string) (*PersonResponse, error) {
	// check value is not empty
	if value == "" {
		return nil, errors.New("Value missing.")
	}

	// Construct URL
	baseUrl, err := url.Parse(p.url)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("apiKey", p.apiKey)
	params.Add(method, value)

	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pr PersonResponse
	err = json.Unmarshal(body, &pr)
	if err != nil {
		return nil, err
	}

	return &pr, nil

}

func (p *PersonAPI) GetByEmail(email string) (*PersonResponse, error) {
	return getPersonBy(p, "email", email)
}

func (p *PersonAPI) GetByTwitter(twitter string) (*PersonResponse, error) {
	return getPersonBy(p, "twitter", twitter)
}

func (p *PersonAPI) GetByPhone(phone string) (*PersonResponse, error) {
	return getPersonBy(p, "phone", phone)
}

func (p *PersonAPI) GetByFacebookUsername(facebookUsername string) (*PersonResponse, error) {
	return getPersonBy(p, "facebookUsername", facebookUsername)
}

func (p *PersonAPI) GetByFacebookID(facebookID string) (*PersonResponse, error) {
	return getPersonBy(p, "facebookId", facebookID)
}
