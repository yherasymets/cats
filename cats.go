package cats

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
)

const (
	apiURL   = "https://catfact.ninja/breeds"
	filePath = "./bin/out.json"
)

// Cat entity
type Cat struct {
	Breed   string `json:"breed"`
	Country string `json:"country"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}

type Cats struct {
	Data []Cat `json:"data"`
}

// GetCetBreeds is function for retrive a list of cat breeds via API
func GetCetBreeds() (*Cats, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response body: %w", err)
	}
	defer resp.Body.Close()

	var cats *Cats
	if err := json.Unmarshal(rawBody, &cats); err != nil {
		return nil, fmt.Errorf("can't unmarshal JSON: %w", err)
	}
	return cats, nil
}

// GroupBreedsByCountry is function for grouping breeds by their country
func GroupBreedsByCountry(cats *Cats) map[string][]string {
	groupedBreeds := make(map[string][]string)
	for _, cat := range cats.Data {
		groupedBreeds[cat.Country] = append(groupedBreeds[cat.Country], cat.Breed)
	}
	return groupedBreeds
}

// SortBreedsByLength is function for sort breeds by length
func SortBreedsByLength(groupedBreeds map[string][]string) {
	for country := range groupedBreeds {
		sort.Slice(groupedBreeds[country], func(i, j int) bool {
			return len(groupedBreeds[country][i]) < len(groupedBreeds[country][j])
		})
	}
}

// WriteToJSON is function for write data into JSON file
func WriteToJSON(groupedBreeds map[string][]string) error {
	data, err := json.MarshalIndent(groupedBreeds, "", " ")
	if err != nil {
		return fmt.Errorf("can't marshal data to JSON: %w", err)
	}
	return os.WriteFile(filePath, data, 0644)
}
