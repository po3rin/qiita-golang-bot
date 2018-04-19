package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Item for decoding JSON
type Item struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
}

// RawGet for creating cURL
type RawGet struct {
	endpoint string
	time     string
	tag      string
}

func (r *RawGet) getQiitaItem() *[]Item {
	u, err := url.Parse(r.endpoint)
	if err != nil {
		log.Fatal("エンドポイントがパースできませんでした: ", err)
	}

	q := u.Query()
	q.Set("query", "tag:"+r.tag+" created:>"+r.time)
	q.Set("page", "1")
	q.Set("per_page", "10")
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")

	res, err := new(http.Client).Do(req)
	if err != nil {
		log.Fatal("リクエストに失敗しました: ", err)
	}
	defer res.Body.Close()

	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var items []Item
	if err := json.Unmarshal(contents, &items); err != nil {
		log.Fatal("JSONデコードに失敗しました: ", err)
	}
	return &items
}
