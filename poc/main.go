package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	url := "https://locations.patch3s.dev/api/random"
	random, err := fetchLocations(url)

	// Check for 3rd party API errors
	if err != nil {
		errMess := "Forecast unavailable"
		http.Error(w, errMess, http.StatusInternalServerError)
		return
	}

	// Check for nil Locations
	if random == nil || len(random.Locations) == 0 {
		errMess := "No locations available"
		http.Error(w, errMess, http.StatusInternalServerError)
		return
	}


	fmt.Println("main.go@18 waldo: forecast", forecast)
}

func main() {
	//register generic handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Prevent any additional url pathing
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		handler(w, r)
	})
}

func fetchForecast(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
// fetchLocations fetches a random location from the 3rd party API
func fetchLocations(url string) (*Random, error) {
	body, err := fetchUrl(url)
	if err != nil {
		return nil, err
	}

	var result Random
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// fetchUrl Helper function to make requests to 3rd party APIs
func fetchUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	// TODO: Consider reading response error code from 3rd party API and
	// converting to a more friendly error message
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Random represents the response from 3rd party /random
type Random struct {
	Locations []struct {
		Lat  float64 `json:"latitude"`
		Lon  float64 `json:"longitude"`
		Name string  `json:"name"`
	} `json:"locations"`
}
