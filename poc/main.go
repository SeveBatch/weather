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
	forecast, _ := fetchForecast(url)

	fmt.Println("main.go@18 waldo: forecast", forecast)
}

func main() {
	//register generic handler
	http.HandleFunc("/", handler)
	log.Println("Server starting on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func fetchForecast(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
