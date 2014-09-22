package main

import (
  "log"
  "os"
  "time"
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

  for ;; {
    log.Printf("Update complete! Sleeping for 60s...")
    time.Sleep(time.Second * 60)

    log.Printf("Checking for updated cookbooks...")
    updates, err := collection.Update()
    if err != nil {
      log.Printf("WARNING: Error occurred during update: '%s'", err)
      continue
    }
    for _, cb := range updates.Cookbooks {
      for _, version := range cb.Versions {
        log.Printf("Found new entry, cookbook: '%s', version: '%s'",
                   cb.Name,
                   version)
      }
    }
  }
}
