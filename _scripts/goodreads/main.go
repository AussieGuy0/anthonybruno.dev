package main

import (
	"log"
)

func main() {
	readBooks, err := GetReadBooks(51361759)
	if err != nil {
		log.Fatal(err)
	}
	siteBaseDir := "../../_read_books"
	err = WriteReviews(readBooks, siteBaseDir)
	if err != nil {
		log.Fatal(err)
	}
}
