package slack

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/slack-go/slack"
)

// Fetch uses the Slack API to find messages and its replies
// it calls resultCallback once per message found
func Fetch(channelName string, resultCallback func(Message)) {
	log.Println("fetching slack api")
	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true))

	log.Println("joining channel")
	channel := findAndJoinChannel(api, channelName)

	if channel == nil {
		fmt.Println("Channel not found :'(")
		return
	}

	startDate := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	log.Println("fetching channel history")
	messages, channelHistoryErr := getChannelHistory(api, channel, startDate, endDate)

	if channelHistoryErr != nil {
		fmt.Printf("error getting channel history: %s\n", channelHistoryErr)
	}

	memoizedUserInfo := createUserMemoizer(api)
	for _, msg := range messages {

		user, err := memoizedUserInfo(msg.User)

		if err == nil {
			log.Println("fetching message info")
			permalink, _ := api.GetPermalink(&slack.PermalinkParameters{
				Ts:      msg.Timestamp,
				Channel: channel.ID,
			})

			resultCallback(NewMessage(
				msg.Text,
				blockchainsForMessage(api, channel, msg.Timestamp),
				user.Name,
				repliesForMessage(api, channel, msg, memoizedUserInfo),
				permalink,
			))
		}

	}
}
