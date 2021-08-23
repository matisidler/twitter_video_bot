package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Credentials struct {
	AccessToken       string
	ConsumerKey       string
	ConsumerSecret    string
	AccessTokenSecret string
}

func getClient(creds *Credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig("ConsumerKey", "ConsumerSecret")
	token := oauth1.NewToken("AccessToken", "AccessTokenSecret")

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	return client, nil
}

var i = 1

func Testing() {
	for {
		client, err := getClient(&Credentials{})
		if err != nil {
			fmt.Println(err)
		}
		params := &twitter.MentionTimelineParams{Count: 10}

		tweet, _, err := client.Timelines.MentionTimeline(params)
		if len(tweet) == 0 {
			fmt.Println("no mentions yet")
		}

		for _, tw := range tweet {
			rpstr := tw.InReplyToStatusIDStr
			rpint := tw.InReplyToStatusID
			idint := tw.ID
			twshow, _, err := client.Statuses.Show(rpint, nil)
			if err != nil {
				fmt.Println("not a reply")
				break
			}
			name := tw.User.ScreenName
			params := &twitter.StatusUpdateParams{InReplyToStatusID: idint}

			if len(twshow.Entities.Media) == 0 {
				_, _, err = client.Statuses.Update("@"+name+" ¿Sos estúpido? Esto no es un video", params)
				if err != nil {
					fmt.Println("Last step: ", err)
				}
				break
			} else if len(twshow.Entities.Media) > 0 {
				for _, i := range twshow.Entities.Media {
					if strings.Contains(i.URLEntity.ExpandedURL, "photo") {
						_, _, err = client.Statuses.Update("@"+name+" Esto es una foto, pavo.", params)
						if err != nil {
							fmt.Println("Last step: ", err)
						}
						break
					} else if strings.Contains(i.URLEntity.ExpandedURL, "video") {
						_, _, err = client.Statuses.Update("@"+name+` Acá tenés tu video. Disfrutalo! // Here is your download link! Enjoy it: https://ssstwitter.com/i/status/`+rpstr, params)

						if err != nil {
							fmt.Println("Last step: ", err)
						}
						break

					}
				}
			}

		}
		time.Sleep(50 * time.Second)
	}
}

func main() {
	for {
		Testing()
		time.Sleep(50 * time.Second)
	}
}
