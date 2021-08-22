package funcs

import (
	"fmt"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/gofiber/fiber/v2"
)

type Credentials struct {
	AccessToken       string //1428812771819466757-FKzRVhyjBaj9jcNPluVKPPjXs4lOec
	ConsumerKey       string //"b9RrONp8tqd7WmI2p3Pq8tkGY"
	ConsumerSecret    string //j9cyDU7xcdzaxPQ9saqS9DrM7Aq4OyyunQLJxozlt1MHNv3vYd
	AccessTokenSecret string //0M98FOe1T98zwKZnkNJ5FjGDQ2KyvDpODVTVhpFd9feja
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

func Testing(f *fiber.Ctx) error {
	for {
		client, err := getClient(&Credentials{})
		if err != nil {
			return err
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
