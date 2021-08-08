package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/rohithv1997/SentimentAnalysis-Go/applicationConfig"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ApiEndpoint(searchTerm ...string) {

	consumerKey := applicationConfig.Instance.GetValue(ConsumerKey)
	consumerSecret := applicationConfig.Instance.GetValue(ConsumerSecret)
	accessToken := applicationConfig.Instance.GetValue(AccessToken)
	accessSecret := applicationConfig.Instance.GetValue(AccessSecret)

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter Client
	client := twitter.NewClient(httpClient)

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		//	fmt.Printf("Text: %s\r\n",tweet.Text)
	}
	fmt.Println("Starting Stream...")

	// FILTER
	filterParams := &twitter.StreamFilterParams{
		Track:         searchTerm,
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}

	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "gopher",
	})

	//for _, s := range search.Statuses {
	//	fmt.Println(s.s)
	//}

	serialised, _ := json.Marshal(search.Statuses[0])
	fmt.Printf("%s\n\n\n****\n\n\n",serialised)

	serialised, _ = json.Marshal(search.Metadata)
	fmt.Printf("METADATA: \n\n\n%s\n", serialised)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
}
