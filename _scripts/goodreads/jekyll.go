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

func WriteReviews(reviews []Review, writeFolder string) error {
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

func shouldSkipReview(review *Review) bool {
	// If readAt is an empty string, the books was not read and we can safely skip it
	if len(review.ReadAt) == 0 {
		log.Println("Skipping " + review.Book.Title + " as it has not readAt date")
		return true
	}

	return false
}

func writeReview(review *Review, path string) error {
	str := generateReviewString(review)
	filename, err := generateFilename(review)
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

func generateFilename(review *Review) (string, error) {
	readAt, err := review.ReadAtTime()
	if err != nil {
		return "", err
	}
	title := strings.ReplaceAll(review.Book.Title, " ", "-")
	return readAt.Format(dateFormat) + "-" + title + ".md", nil
}

func generateReviewString(review *Review) string {
	var sb strings.Builder
	content := strings.TrimSpace(review.Body)
	sb.WriteString("---\n")
	writeKey(&sb, "rating", strconv.Itoa(review.Rating))
	writeKey(&sb, "title", "\""+review.Book.Title+"\"")
	writeKey(&sb, "link", review.Book.Link)
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
