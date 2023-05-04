package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type EarthImageData []struct {
	Identifier string `json:"identifier"`
	Caption    string `json:"caption"`
	Image      string `json:"image"`
	Version    string `json:"version"`
}

func main() {
	apiKey := os.Getenv("API_KEY")
	baseUrl := "https://api.nasa.gov/EPIC/api/natural/%s?api_key=" + apiKey
	getEarthImageData(baseUrl, "2023-01-23")
}

func getEarthImageData(url, date string) (EarthImageData, error) {
	endpoint := "date/" + date
	earthDataUri := fmt.Sprintf(url, endpoint)
	fmt.Println(earthDataUri)
	res, err := http.Get(earthDataUri)
	if err != nil {
		fmt.Println("Error while making the request", err)
		return nil, err
	}

	defer res.Body.Close()

	var imageData EarthImageData

	if err := json.NewDecoder(res.Body).Decode(&imageData); err != nil {
		log.Println(err)
	}

	fmt.Println(imageData)

	return imageData, nil
}
