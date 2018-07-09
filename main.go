package main

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/po3rin/qiitter/oauth"
	"github.com/po3rin/qiitter/qiita"
	"golang.org/x/sync/errgroup"
)

var hash = os.Getenv("HASH_TAG")

func post() error {
	var c = qiita.Client{
		Endpoint: "https://qiita.com/api/v2/items",
		Time:     time.Now().Format("2006-01-02"),
		Tag:      os.Getenv("TARGET_TAG"),
	}

	ch1 := make(chan *twitter.Client)
	ch2 := make(chan *[]qiita.Item)

	go oauth.Client(ch1)
	go c.GetQiitaItems(ch2)

	client := <-ch1
	items := <-ch2

	boundary := time.Now().Add(time.Duration(-1) * time.Hour).Unix()

	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, i := range *items {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return errors.New("Error Post")
			default:
				t, err := time.Parse("2006-01-02T15:04:05+09:00", i.CreatedAt)
				if err != nil {
					cancel()
				}

				created := t.Add(time.Duration(-9) * time.Hour)
				createdString := t.Format("01/02 15:04")
				createdUnix := created.Unix()

				if createdUnix > boundary {
					post := createdString + "に投稿されました\n" + i.Title + "\n" + hash + "\n" + i.URL
					_, _, err := client.Statuses.Update(post, nil)
					if err != nil {
						cancel()
					}
				}
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func main() {
	lambda.Start(post)
}
