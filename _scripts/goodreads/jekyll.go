package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	dateFormat = "2006-01-02"
)

func WriteReviews(reviews []ReadBook, writeFolder string) error {
	for _, review := range reviews {
		if shouldSkipReview(&review) {
			continue
		}
		err := writeReview(&review, writeFolder)
		if err != nil {
			return err
		}
	}
	return nil
}

func shouldSkipReview(readBook *ReadBook) bool {
	// If readAt is an empty string, the books was not read and we can safely skip it
	if len(readBook.ReadAt) == 0 && len(readBook.Added) == 0 {
		log.Println("Skipping " + readBook.Title + " as it has no readAt date")
		return true
	}

	return false
}

func writeReview(readBook *ReadBook, path string) error {
	str := generateReviewString(readBook)
	filename, err := generateFilename(readBook)
	if err != nil {
		return err
	}

	filePath := path + "/" + filename
	log.Println("Writing to " + filePath)
	err = ioutil.WriteFile(filePath, []byte(str), 0644)
	if err != nil {
		return err
	}
	return nil
}

func generateFilename(readBook *ReadBook) (string, error) {
	readAt, err := readBook.ReadAtTime()
	if err != nil {
		return "", err
	}
	title := strings.ReplaceAll(readBook.Title, " ", "-")
	title = strings.ReplaceAll(title, "/", "-")
	return readAt.Format(dateFormat) + "-" + title + ".md", nil
}

func generateReviewString(readBook *ReadBook) string {
	var sb strings.Builder
	content := strings.TrimSpace(readBook.Review)
	sb.WriteString("---\n")
	writeKey(&sb, "rating", strconv.Itoa(readBook.Rating))
	writeKey(&sb, "title", "\""+readBook.Title+"\"")
	writeKey(&sb, "link", readBook.Link)
	writeKey(&sb, "has_content", strconv.FormatBool(len(content) > 0))
	writeKey(&sb, "layout", "book")
	sb.WriteString("---\n")
	sb.WriteString(content)
	return sb.String()
}

func writeKey(sb *strings.Builder, key string, value string) {
	sb.WriteString(key)
	sb.WriteString(": ")
	sb.WriteString(value)
	sb.WriteString("\n")
}
