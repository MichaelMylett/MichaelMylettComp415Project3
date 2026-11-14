package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const APIServerAddress = "https://api.restcountries.com/countries/v5?region="

// All of these structs for reading the API are AI generated with minimal changes made by myself

type APIResponse struct {
	Data ResponseData `json:"data"`
}
type ResponseData struct {
	Meta    Meta      `json:"meta"`
	Objects []Country `json:"objects"`
}
type Meta struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
type Country struct { //2 extra relevant pieces of data: population and currencies
	Names      Names      `json:"names"`
	Region     string     `json:"region"`
	Languages  []Language `json:"languages"`
	Currencies []Currency `json:"currencies"`
	Population int64      `json:"population"`
	Flag       Flag       `json:"flag"`
	Landlocked bool       `json:"landlocked"`
}
type Names struct {
	Common   string                `json:"common"`
	Official string                `json:"official"`
	Native   map[string]NativeName `json:"native,omitempty"`
}
type NativeName struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}
type Language struct {
	Name  string `json:"name"`
	BCP47 string `json:"bcp47"`
}
type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
type Flag struct {
	Emoji       string `json:"emoji"`
	URLSVG      string `json:"url_svg"`
	URLPNG      string `json:"url_png"`
	Description string `json:"description,omitempty"`
}

func getCountryData(region string) ([]Country, error) { // originally also had "countryCode string" in parameters but I removed it because it felt unnecessary
	var result []Country // a lot of this uses the rest demo as a template
	apiClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	limiter := "&limit=100"
	fullAddress := fmt.Sprintf("%s%s%s", APIServerAddress, region, limiter) //had to add limiter because it originally capped the list length, noticed because I could not find the US
	request, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet, fullAddress, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", APIKey) // defined as a public variable in a file APIKeyContainer, which I will not attach to the project submission for safety
	response, err := apiClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	defer response.Body.Close()
	var jsonResponse = APIResponse{}
	err = json.NewDecoder(response.Body).Decode(&jsonResponse)
	if err != nil && err != io.EOF {
		print("Error the actual response body was: ", response.Body, err)
		return nil, err
	}

	result = jsonResponse.Data.Objects
	return result, nil
}
