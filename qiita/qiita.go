package qiita

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Client - create cURL
type Client struct {
	Endpoint string
	Time     string
	Tag      string
}

// Item - Qiita Post
type Item struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
}

func makeQuery(u *url.URL, c *Client) (q string) {
	query := u.Query()
	query.Set("query", "tag:"+c.Tag+" created:>"+c.Time)
	query.Set("page", "1")
	query.Set("per_page", "8")
	u.RawQuery = query.Encode()
	return u.String()
}

// GetQiitaItems qiita API client
func (c *Client) GetQiitaItems(ch chan<- *[]Item) {
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		log.Fatal("エンドポイントがパースできませんでした: ", err)
	}

	q := makeQuery(u, c)
	req, _ := http.NewRequest("GET", q, nil)
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
	ch <- &items
}
