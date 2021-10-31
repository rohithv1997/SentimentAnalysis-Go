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

func SearchApi(searchTerm string) {
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

	search, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: searchTerm,
	})

	fmt.Println("Length = ", len(search.Statuses))

	//serialised, _ := json.Marshal(search.Statuses[0])
	//fmt.Printf("%s\n\n\n****\n\n\n", serialised)

	// fmt.Println("Text - ", search.Statuses[0])

	serialised, _ := json.Marshal(search.Statuses[0])
	fmt.Printf("METADATA: \n\n\n%s\n", serialised)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
}
