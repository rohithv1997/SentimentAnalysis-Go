package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func StreamApi(searchTerm string, successAction func(string)) {
	if getConfigInstance().consumerKey == "" ||
		getConfigInstance().consumerSecret == "" ||
		getConfigInstance().accessToken == "" ||
		getConfigInstance().accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(getConfigInstance().consumerKey, getConfigInstance().consumerSecret)
	token := oauth1.NewToken(getConfigInstance().accessToken, getConfigInstance().accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter Client
	client := twitter.NewClient(httpClient)

	params := &twitter.StreamFilterParams{
		Track:         []string{searchTerm},
		StallWarnings: twitter.Bool(true),
		Language:      []string{"en"},
	}
	stream, err := client.Streams.Filter(params)
	if err != nil {
		log.Fatal(err)
	}

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Printf(tweet.Text)
		successAction(tweet.Text)
		//serialised, _ := json.Marshal(tweet.Text)
		//fmt.Printf("%s\n\n\n****\n\n\n", serialised)
	}
	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
}
