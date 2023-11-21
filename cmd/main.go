package main

import (
	"log"

	"github.com/yherasymets/cats"
)

func main() {
	// Getting a list of cat breeds via API
	catsData, err := cats.GetCatBreeds()
	if err != nil {
		log.Printf("Failed to retrive the list of cat breeds: %v", err)
		return
	}

	// Grouping of breeds by country of origin
	breeds := cats.GroupBreedsByCountry(catsData)

	// Sorting breeds by length
	cats.SortBreedsByLength(breeds)

	// Writing data to JSON file
	if err := cats.WriteToJSON(breeds); err != nil {
		log.Printf("Error writing to JSON file: %v", err)
		return
	}
	log.Println("Data was recorded to ./bin/out.json")
}
