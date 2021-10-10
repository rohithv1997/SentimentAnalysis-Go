package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ApiEndpoint(searchTerm string) {
	if GetTwitterConfigInstance().consumerKey == "" ||
		GetTwitterConfigInstance().consumerSecret == "" ||
		GetTwitterConfigInstance().accessToken == "" ||
		GetTwitterConfigInstance().accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(GetTwitterConfigInstance().consumerKey, GetTwitterConfigInstance().consumerSecret)
	token := oauth1.NewToken(GetTwitterConfigInstance().accessToken, GetTwitterConfigInstance().accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter Client
	client := twitter.NewClient(httpClient)

	search, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: searchTerm,
	})

	fmt.Println("Length = ", len(search.Statuses))

	serialised, _ := json.Marshal(search.Statuses[0])
	fmt.Printf("%s\n\n\n****\n\n\n", serialised)

	fmt.Println("Text - ", search.Statuses[0].Text)

	serialised, _ = json.Marshal(search.Metadata)
	fmt.Printf("METADATA: \n\n\n%s\n", serialised)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
}
