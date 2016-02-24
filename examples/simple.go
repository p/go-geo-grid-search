package main

import (
	"fmt"
	"github.com/p/go-geo-grid-search"
)

type Store struct {
	name string
	lat  float64
	lng  float64
}

func (s Store) Lat() float64 {
	return s.lat
}

func (s Store) Lng() float64 {
	return s.lng
}

func main() {
	searcher := ggsearch.NewSearcher(100, 200)
	searcher.AddLocatable(Store{"HQ", 40, -70})
	searcher.AddLocatable(Store{"Riverside", 40.2, -70.1})
	searcher.AddLocatable(Store{"Sunset", 39.4, -69.5})

	fmt.Printf("Simple search:\n")
	results := searcher.Search(nil, 40, -70, 10)
	for _, result := range results {
		fmt.Printf("%s @ %f miles\n", result.Locatable().(Store).name, result.Distance())
	}

	fmt.Printf("\nName filtering:\n")
	filter := func(locatable ggsearch.Locatable) bool {
		return len(locatable.(Store).name) > 3
	}
	results = searcher.Search((*ggsearch.Filter)(&filter), 40, -70, 10)
	for _, result := range results {
		fmt.Printf("%s @ %f miles\n", result.Locatable().(Store).name, result.Distance())
	}

	fmt.Printf("\nLimiting results:\n")
	results = searcher.Search(nil, 40, -70, 2)
	for _, result := range results {
		fmt.Printf("%s @ %f miles\n", result.Locatable().(Store).name, result.Distance())
	}
}
