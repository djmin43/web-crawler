package main

import (
	"encoding/json"
	"github.com/gocolly/colly/v2"
	"log"
	"os"
	"time"
)

func main() {

	fName := "pokemon.json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	domain := "pokemongo.inven.co.kr"
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.MaxDepth(1),
	)

	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: domain + "/*",
		// Set a delay between requests to these domains
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})

	pokemons := make([]string, 0, 200)

	// Find and visit all links
	c.OnHTML(".pokemonname", func(e *colly.HTMLElement) {
		//fmt.Println(e.Text)
		pokemons = append(pokemons, e.Text)
	})

	c.Visit("https://pokemongo.inven.co.kr/dataninfo/pokemon/")

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(pokemons)
}
