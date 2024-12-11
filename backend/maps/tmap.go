package maps

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

var appKeyHeader string = "appKey"
var appKey string = os.Getenv("TMAP_API_KEY")

var pedestrianURL string = "https://apis.openapi.sk.com/tmap/routes/pedestrian?version=1"

func GetPedestrianTimeAsMinutes(startLongitude, startLatitude, endLongitude, endLatitude float64, startAddress, endAddress string) (float64, error) {
	requestBody := map[string]string{
		"startX":    strconv.FormatFloat(startLongitude, 'f', -1, 64),
		"startY":    strconv.FormatFloat(startLatitude, 'f', -1, 64),
		"endX":      strconv.FormatFloat(endLongitude, 'f', -1, 64),
		"endY":      strconv.FormatFloat(endLatitude, 'f', -1, 64),
		"startName": startAddress,
		"endName":   endAddress,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", pedestrianURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(appKeyHeader, appKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Response", resp.StatusCode)
		return 0, errors.New("failed to get pedestrian time")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, err
	}

	features := response["features"].([]interface{})
	if len(features) > 0 {
		properties := features[0].(map[string]interface{})["properties"].(map[string]interface{})
		totalTime := properties["totalTime"].(float64)
		return totalTime / 60, nil
	}

	return 0, errors.New("no features found")
}
