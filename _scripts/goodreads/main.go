package main

import (
	"log"
	"os"
)

func main() {
	key := os.Getenv("GOODREADS_KEY")
    if key == "" {
        log.Fatal("Requries env property GOODREADS_KEY")
    }
	readBooks, err := GetReviews(51361759, key)
	if err != nil {
		log.Fatal(err)
	}
    siteBaseDir := "../../_read_books"
    err = WriteReviews(readBooks, siteBaseDir)
    if (err != nil) {
		log.Fatal(err)
    }
}
