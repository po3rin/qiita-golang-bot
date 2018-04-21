package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/po3rin/qiitter/oauth"
	"github.com/po3rin/qiitter/qiita"
)

// Post post twitter
func post() {
	var r = qiita.RawGet{
		Endpoint: "https://qiita.com/api/v2/items",
		Time:     time.Now().Format("2006-01-02"),
		Tag:      "Go",
	}

	ch1 := make(chan *twitter.Client)
	ch2 := make(chan *[]qiita.Item)

	go oauth.Client(ch1)
	go r.GetQiitaItem(ch2)

	client := <-ch1
	items := <-ch2

	boundary := time.Now().Add(time.Duration(-1) * time.Hour).Unix()
	var wg sync.WaitGroup
	for _, i := range *items {
		wg.Add(1)
		go func(i qiita.Item) {
			defer wg.Done()
			t, _ := time.Parse("2006-01-02T15:04:05+09:00", i.CreatedAt)
			created := t.Add(time.Duration(-9) * time.Hour).Unix()
			if created > boundary {
				post := i.Title + i.URL
				_, _, err := client.Statuses.Update(post, nil)
				if err != nil {
					fmt.Println("投稿に失敗しました: ", err)
				}
			}
		}(i)
	}
	wg.Wait()
}

func main() {
	lambda.Start(post)
}
