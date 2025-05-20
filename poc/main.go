package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// TODO: Update code for high load and concurrency by using a pool or queue of available
// connections or go channels or routines to handle multiple requests

// Handle request to "/"
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

	baseCoordUrl := "https://api.weather.gov/points/"

	// Single forecast object used but could be updated to a slice of forecasts
	// to support multiple locations
	var forecast *Forecast
	// Parse locations for coordinates
	for _, f := range random.Locations {
		coords := baseCoordUrl + fmt.Sprintf("%f,%f", f.Lat, f.Lon)
		log.Printf("Constructing coordinates URL: %s", coords)
		forecast, err = FetchForecast(coords)
		if err != nil {
			errMess := "Forecast unavailable"
			http.Error(w, errMess, http.StatusInternalServerError)
		}
	}

	// Successfully fetched and generated forecast
	if forecast != nil {
		forecast.Name = random.Locations[0].Name
		log.Printf("Writing response for %s", forecast.Name)
		jsonResponse := []byte(forecast.Detailed)
		if err != nil {
			http.Error(w, "Error generating response", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
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

	//TODO change back to 5000
	log.Println("Server starting on port 5001...")
	log.Fatal(http.ListenAndServe(":5001", nil))
}

// FetchForecast fetches the forecast for a given set of coordinates
func FetchForecast(url string) (*Forecast, error) {
	var forecast Forecast
	body, err := fetchUrl(url)
	if err != nil {
		return nil, err
	}

	log.Printf("Unmarshalling forecast url from weather.gov")
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		return nil, err
	}

	body, err = fetchUrl(forecast.Properties.Url)
	if err != nil {
		return nil, err
	}

	var points Points
	log.Printf("Unmarshalling forecast points data from weather.gov")
	err = json.Unmarshal(body, &points)
	if err != nil {
		return nil, err
	}

	// The first period represents the current forecast
	// TODO: This could be extended to support other time frames returned from weather.gov
	forecast.Detailed = points.Properties.Periods[0].Detailed
	forecast.Short = points.Properties.Periods[0].Short

	return &forecast, nil
}

// fetchLocations fetches a random location from the 3rd party API
func fetchLocations(url string) (*Random, error) {
	log.Printf("Fetching Locations from patch3s.dev")
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
	log.Printf("Fetching URL: %s", url)
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

// Models
// Forecast represents the forecast data for a locations forecasts
type Forecast struct {
	Detailed   string
	Name       string
	Properties struct {
		Url string `json:"forecast"`
	} `json:"properties"`
	Short string
}

// Points represents the response from weather.gov
type Points struct {
	Properties struct {
		Periods []struct {
			Detailed string `json:"detailedForecast"`
			Short    string `json:"shortForecast"`
		} `json:"periods"`
	} `json:"properties"`
}

// Random represents the response from 3rd party /random
type Random struct {
	Locations []struct {
		Lat  float64 `json:"latitude"`
		Lon  float64 `json:"longitude"`
		Name string  `json:"name"`
	} `json:"locations"`
}
