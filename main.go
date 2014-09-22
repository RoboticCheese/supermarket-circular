package main

import (
  "log"
  "os"
  "github.com/RoboticCheese/supermarket-circular/cookbook_collection"
)

const (
  API_ENDPOINT = "https://supermarket.getchef.com"
  UNIVERSE = "universe.json"
)

func main() {
  log.SetOutput(os.Stdout)

  url := API_ENDPOINT + "/" + UNIVERSE
  log.Printf("Starting up with URL '%s'...", url)

  log.Printf("Initializing current state of cookbook universe...")
  collection, err := new(cookbook_collection.CookbookCollection).NewFromUniverse(url)
  if err != nil {
    log.Fatal(err)
  }
  for _, cb := range collection.Cookbooks {
    log.Printf("Read cookbook: '%s', versions: '%s'", cb.Name, cb.Versions)
  }
  log.Printf("Initialization complete!")
}
