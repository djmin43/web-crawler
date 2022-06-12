package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"os"
	"strings"
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

	type Pokemon struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	pokemons := make([]Pokemon, 0, 200)

	// Find and visit all links
	c.OnHTML(".pokemonname", func(e *colly.HTMLElement) {
		s := strings.Split(e.Text, ".")
		name := s[1]
		p := Pokemon{
			Id:   s[0],
			Name: s[1][0 : len(name)-1],
		}
		pokemons = append(pokemons, p)
	})

	c.Visit("https://pokemongo.inven.co.kr/dataninfo/pokemon/")

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	fmt.Println(pokemons)
	enc.Encode(pokemons)

}
