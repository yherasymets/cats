package main

import (
	"log"

	"github.com/yherasymets/cats"
)

func main() {
	// Getting a list of cat breeds via API
	catsData, err := cats.GetCetBreeds()
	if err != nil {
		log.Printf("Failed to retrive the list of cat breeds: %v", err)
		return
	}

	// Grouping of breeds by country of origin
	groupedBreeds := cats.GroupBreedsByCountry(catsData)

	// Sorting breeds by length
	cats.SortBreedsByLength(groupedBreeds)

	// Writing data to JSON file
	if err := cats.WriteToJSON(groupedBreeds); err != nil {
		log.Printf("Error writing to JSON file: %v", err)
		return
	}
	log.Println("Data was recorded to out.json")
}
