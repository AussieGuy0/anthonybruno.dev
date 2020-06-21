package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
    "errors"
    "time"
    "log"
)

var goodReadsUrl = "https://www.goodreads.com"

type ReadBooks struct {
	XMLName xml.Name `xml:"GoodreadsResponse"`
	Reviews *Reviews `xml:"reviews"`
}

type Reviews struct {
	XMLName    xml.Name `xml:"reviews"`
	ReviewList []Review `xml:"review"`
}

type Review struct {
	XMLName xml.Name `xml:"review"`
	Book    *Book    `xml:"book"`
	Rating  int      `xml:"rating"`
	ReadAt  string   `xml:"read_at"`
	Body string   `xml:"body"`
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	Title   string   `xml:"title_without_series"`
	Link    string   `xml:"link"`
}

func GetReviews(userId int, key string) ([]Review, error) {
	res, err := makeReadBooksRequest(userId, key)
	if err != nil {
		return nil, err
	}
	readBooks, err := parseReadBooks(res)
	if err != nil {
		return nil, err
	}
	return readBooks.Reviews.ReviewList, nil
}

func (review Review) ReadAtTime() (time.Time, error) {
   return time.Parse(time.RubyDate, review.ReadAt)
}

func makeReadBooksRequest(userId int, key string) (*http.Response, error) {
	url, err := url.Parse(goodReadsUrl + "/review/list/" + strconv.Itoa(userId) + ".xml")
	if err != nil {
		return nil, err
	}
	q := url.Query()
	q.Add("key", key)
	q.Add("shelf", "read")
	q.Add("v", "2")
    q.Add("per_page", "200") //TODO: Good enough for now, will need multiple requests in future
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

func parseReadBooks(res *http.Response) (*ReadBooks, error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var parsed ReadBooks
	err = xml.Unmarshal(body, &parsed)
	if err != nil {
		return nil, err
	}
    return &parsed, nil
}
