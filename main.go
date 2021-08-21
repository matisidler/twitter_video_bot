package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//API KEY CFNuJtksb8Iis3I5ePPIaCXRy
//API SECRET KEY WcS28yFOZDTw7iu7U2grl4DXRuDRdVj2d4qleVGImFcNHEXp3T
//BEARER TOKEN AAAAAAAAAAAAAAAAAAAAAHyHSwEAAAAACM9COtz4FQGt7rcpSJajPgk4lkQ%3DBCjmSOmWzu3x8c0TyxY0e4kCbfAadsOsnMMxvTZUsITVhWasG7

// other imports

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type Credentials struct {
	AccessToken       string //1428812771819466757-FKzRVhyjBaj9jcNPluVKPPjXs4lOec
	ConsumerKey       string //"b9RrONp8tqd7WmI2p3Pq8tkGY"
	ConsumerSecret    string //j9cyDU7xcdzaxPQ9saqS9DrM7Aq4OyyunQLJxozlt1MHNv3vYd
	AccessTokenSecret string //0M98FOe1T98zwKZnkNJ5FjGDQ2KyvDpODVTVhpFd9feja
}

type ReplyID struct {
	InReplyToStatusID string `json:"in_reply_to_status_id_str"`
}
type StatusUpdateParams struct {
}

// getClient is a helper function that will return a twitter client
// that we can subsequently use to send tweets, or to stream new tweets
// this will take in a pointer to a Credential struct which will contain
// everything needed to authenticate and return a pointer to a twitter Client
// or an error
func getClient(creds *Credentials) (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig("b9RrONp8tqd7WmI2p3Pq8tkGY", "j9cyDU7xcdzaxPQ9saqS9DrM7Aq4OyyunQLJxozlt1MHNv3vYd")
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken("1428812771819466757-FKzRVhyjBaj9jcNPluVKPPjXs4lOec", "0M98FOe1T98zwKZnkNJ5FjGDQ2KyvDpODVTVhpFd9feja")

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	/* log.Printf("User's ACCOUNT:\n%+v\n", user) */
	return client, nil
}

func newTrue() *bool {
	b := true
	return &b
}

func testing() error {
	client, err := getClient(&Credentials{})
	if err != nil {
		return err
	}
	params := &twitter.MentionTimelineParams{Count: 1, IncludeEntities: newTrue()}

	tweet, _, err := client.Timelines.MentionTimeline(params)
	if len(tweet) == 0 {
		fmt.Println("no mentions yet")
	}

	for _, tw := range tweet {
		rpstr := tw.InReplyToStatusIDStr
		rpint := tw.InReplyToStatusID
		idint := tw.ID
		var ids []int
		ids = append(ids, 1)
		for _, id := range ids {
			if id != int(idint) {
				ids = append(ids, int(idint))
				//show
				twshow, _, err := client.Statuses.Show(rpint, nil)
				/* name := tw.User.ScreenName
				params := &twitter.StatusUpdateParams{InReplyToStatusID: idint} */
				name := tw.User.ScreenName
				params := &twitter.StatusUpdateParams{InReplyToStatusID: idint}
				if len(twshow.Entities.Media) == 0 {
					_, _, err = client.Statuses.Update("@"+name+" ¿Sos estúpido? Esto no es un video", params)
					if err != nil {
						fmt.Println("Last step: ", err)
					}
					break
				}
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
			} else {
				fmt.Println("Duplicated tweet with id: ", idint)
			}
		}

	}
	return nil
}

func main() {
	for {
		testing()
		time.Sleep(5 * time.Second)
	}

	/* 	params := &twitter.MentionTimelineParams{Count: 1, IncludeEntities: newTrue()}

	   	tweet, resp, err := client.Timelines.MentionTimeline(params)
	   		tweet, _, err := client.Statuses.Show(585613041028431872, nil)
	   	 _, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
	   		Query: "@this_vid",
	   	})
	   	 tweet, _, err := client.Statuses.Update("just setting up my twttr", nil)
	   	if err != nil {
	   		fmt.Printf("error: %v\n", err)
	   	}
	   	fmt.Println(tweet)
	   	a := twitter.Entities{}  */

}
