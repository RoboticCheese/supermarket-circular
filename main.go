package main

import (
	"github.com/RoboticCheese/goulash/api_instance"
	"github.com/RoboticCheese/goulash/universe"
	"github.com/RoboticCheese/supermarket-circular/config"
	"log"
	"os"
	"time"
)

func log_cookbooks(u *universe.Universe) {
	for _, cb_data := range u.Cookbooks {
		for _, ver_data := range cb_data.Versions {
			log.Printf("Cached cookbook '%v', version '%v'...",
				cb_data.Name,
				ver_data.Version)
		}
	}
}

func main() {
	log.SetOutput(os.Stdout)

	conf, err := config.New("config.json")
	if err != nil {
		log.Fatalf("Error reading config: '%v'", err)
	}
	log.Printf("Starting up with URL '%v'...", conf.Supermarket.Endpoint)

	api, err := api_instance.New(conf.Supermarket.Endpoint)
	if err != nil {
		log.Fatalf("Error connecting to Supermarket API: '%v'", err)
	}
	log.Printf("Connection established to API at '%v'...", api.Endpoint)

	log.Printf("Initializing current state of cookbook universe...")
	univ, err := universe.New(api)
	if err != nil {
		log.Fatalf("Error initializing universe: %v", err)
	}
	log_cookbooks(univ)

	for {
		log.Print("Update complete! Sleeping for 60s...")
		time.Sleep(time.Second * 60)

		log.Print("Checking for updated cookbooks...")
		pos_diff, _, err := univ.Update()
		if err != nil {
			log.Printf("WARNING: Error during update: '%v'", err)
		} else if pos_diff != nil {
			log_cookbooks(pos_diff)
		} else {
			log.Print("No updates found!")
		}
	}
}
