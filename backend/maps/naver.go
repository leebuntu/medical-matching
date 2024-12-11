package maps

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var keyHeader string = "X-NCP-APIGW-API-KEY-ID"
var secretHeader string = "X-NCP-APIGW-API-KEY"
var apiKeyID string = os.Getenv("NAVER_MAP_API_KEY_ID")
var apiKeySecret string = os.Getenv("NAVER_MAP_API_KEY_SECRET")

var geocodeURL string = "https://naveropenapi.apigw.ntruss.com/map-geocode/v2/geocode"
var directionURL string = "https://naveropenapi.apigw.ntruss.com/map-direction/v1/driving"

type Geocode struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type GeocodeResponse struct {
	Addresses []struct {
		Longitude string `json:"x"`
		Latitude  string `json:"y"`
	} `json:"addresses"`
}

type DirectionResponse struct {
	Route struct {
		Traoptimal []struct {
			Summary struct {
				Duration int `json:"duration"`
			} `json:"summary"`
		} `json:"traoptimal"`
	} `json:"route"`
}

func GetGeocode(address string) (*Geocode, error) {
	params := url.Values{}
	params.Set("query", address)

	req, err := http.NewRequest("GET", geocodeURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(keyHeader, apiKeyID)
	req.Header.Set(secretHeader, apiKeySecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geocodeResp GeocodeResponse
	err = json.Unmarshal(body, &geocodeResp)
	if err != nil {
		return nil, err
	}

	if len(geocodeResp.Addresses) > 0 {
		x := geocodeResp.Addresses[0].Longitude
		y := geocodeResp.Addresses[0].Latitude
		xFloat, _ := strconv.ParseFloat(x, 64)
		yFloat, _ := strconv.ParseFloat(y, 64)
		return &Geocode{Longitude: xFloat, Latitude: yFloat}, nil
	}

	return nil, errors.New("no address found")
}

func GetDrivingTimeAsMinutes(startLongitude, startLatitude, endLongitude, endLatitude float64) (int, error) {
	params := url.Values{}
	params.Set("start", fmt.Sprintf("%f,%f", startLongitude, startLatitude))
	params.Set("goal", fmt.Sprintf("%f,%f", endLongitude, endLatitude))

	req, err := http.NewRequest("GET", directionURL+"?"+params.Encode(), nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set(keyHeader, apiKeyID)
	req.Header.Set(secretHeader, apiKeySecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var directionRes DirectionResponse
	err = json.Unmarshal(body, &directionRes)
	if err != nil {
		return 0, err
	}

	if len(directionRes.Route.Traoptimal) > 0 {
		duration := directionRes.Route.Traoptimal[0].Summary.Duration
		return duration / 1000 / 60, nil
	} else {
		return 0, errors.New("no route found")
	}
}
