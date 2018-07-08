package oauth

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//Client for creating Twitter client
func Client(ch chan<- *twitter.Client) {
	var (
		consumerKey    = os.Getenv("CONSUMER_KEY")
		consumerSecret = os.Getenv("CONSUMER_SECRET")
		accessToken    = os.Getenv("ACCESS_TOKEN")
		accessSecret   = os.Getenv("ACCESS_SECRET")
	)

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	ch <- client
}
