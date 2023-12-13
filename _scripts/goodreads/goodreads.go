package main

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var goodReadsUrl = "https://www.goodreads.com"

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	XMLName   xml.Name   `xml:"channel"`
	ReadBooks []ReadBook `xml:"item"`
}

type ReadBook struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Added   string   `xml:"user_date_added"`
	ReadAt  string   `xml:"user_read_at"`
	Review  string   `xml:"user_review"`
	Rating  int      `xml:"user_rating"`
}

func GetReadBooks(userId int) ([]ReadBook, error) {
	res, err := makeReadBooksRequest(userId)
	if err != nil {
		return nil, err
	}
	rss, err := parseRss(res)
	if err != nil {
		return nil, err
	}
	return rss.Channel.ReadBooks, nil
}

func (readBook ReadBook) ReadAtTime() (time.Time, error) {
	layout := "Mon, 2 Jan 2006 15:04:05 -0700"
	if len(readBook.ReadAt) != 0 {
		return time.Parse(layout, readBook.ReadAt)
	}
	// Fallback to added date if no ReadAt date.
	return time.Parse(layout, readBook.Added)
}

func makeReadBooksRequest(userId int) (*http.Response, error) {
	url, err := url.Parse(goodReadsUrl + "/review/list_rss/" + strconv.Itoa(userId))
	if err != nil {
		return nil, err
	}
	q := url.Query()
	q.Add("shelf", "read")
	url.RawQuery = q.Encode()

	log.Println("Getting read books from Goodreads")
	res, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("Non 200 status code: " + strconv.Itoa(res.StatusCode))
	}
	return res, nil
}

func parseRss(res *http.Response) (*Rss, error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var parsed Rss
	err = xml.Unmarshal(body, &parsed)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}
