package cats

import (
	"reflect"
	"testing"
)

func TestSortBreedsByLength(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string][]string
		expected map[string][]string
	}{
		{
			name: "Happy path 1",
			input: map[string][]string{
				"United Kingdom": {"Bambino", "Chantilly-Tiffany", "California Spangled"},
				"United States":  {"Burmilla", "British Shorthair", "British Semi-longhair"},
			},
			expected: map[string][]string{
				"United Kingdom": {"Bambino", "Chantilly-Tiffany", "California Spangled"},
				"United States":  {"Burmilla", "British Shorthair", "British Semi-longhair"},
			},
		},
		{
			name: "Happy path 2",
			input: map[string][]string{
				"United Kingdom": {"Chantilly-Tiffany", "Bambino", "California Spangled"},
				"United States":  {"British Semi-longhair", "Burmilla", "British Shorthair"},
			},
			expected: map[string][]string{
				"United Kingdom": {"Bambino", "Chantilly-Tiffany", "California Spangled"},
				"United States":  {"Burmilla", "British Shorthair", "British Semi-longhair"},
			},
		},
		{
			name: "Incorrect Breed Order",
			input: map[string][]string{
				"United Kingdom": {"Chantilly-Tiffany", "Bambino", "California Spangled"},
				"United States":  {"Burmilla", "British Semi-longhair", "British Shorthair"},
			},
			expected: map[string][]string{
				"United Kingdom": {"Bambino", "Chantilly-Tiffany", "California Spangled"},
				"United States":  {"Burmilla", "British Shorthair", "British Semi-longhair"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			SortBreedsByLength(testCase.input)
			if !reflect.DeepEqual(testCase.input, testCase.expected) {
				t.Errorf("Expected %v, got %v", testCase.expected, testCase.input)
			}
		})
	}
}

func TestGroupBreedsByCountry(t *testing.T) {
	testCases := []struct {
		name     string
		cats     *Cats
		expected map[string][]string
	}{
		{
			name: "Happy path",
			cats: &Cats{
				Data: []Cat{
					{Country: "United Kingdom", Breed: "Bambino"},
					{Country: "Australia", Breed: "Australian Mist"},
					{Country: "United States", Breed: "Burmilla"},
					{Country: "Brazil", Breed: "Brazilian Shorthair"},
					{Country: "United Kingdom", Breed: "Chantilly-Tiffany"},
					{Country: "France", Breed: "Chausie"},
					{Country: "United States", Breed: "British Shorthair"},
					{Country: "Ethiopia", Breed: "Abyssinian"},
					{Country: "United Kingdom", Breed: "California Spangled"},
					{Country: "United States", Breed: "British Semi-longhair"},
					{Country: "France", Breed: "Chartreux"},
				},
			},
			expected: map[string][]string{
				"Australia":      {"Australian Mist"},
				"Brazil":         {"Brazilian Shorthair"},
				"Ethiopia":       {"Abyssinian"},
				"France":         {"Chausie", "Chartreux"},
				"United Kingdom": {"Bambino", "Chantilly-Tiffany", "California Spangled"},
				"United States":  {"Burmilla", "British Shorthair", "British Semi-longhair"},
			},
		},
		{
			name: "Empty input",
			cats: &Cats{
				Data: []Cat{},
			},
			expected: map[string][]string{},
		},
		{
			name: "Single country",
			cats: &Cats{
				Data: []Cat{
					{Country: "United Kingdom", Breed: "Persian"},
					{Country: "United Kingdom", Breed: "Siamese"},
					{Country: "United Kingdom", Breed: "Maine Coon"},
				},
			},
			expected: map[string][]string{
				"United Kingdom": {"Persian", "Siamese", "Maine Coon"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := GroupBreedsByCountry(testCase.cats)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Expected %v, got %v", testCase.expected, result)
			}
		})
	}
}
