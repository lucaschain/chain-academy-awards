package slack

import (
	"log"
	"os"
	"strconv"
	"strings"
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
		log.Println("Channel not found :'(")
		return
	}

	startDate := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	log.Println("fetching channel history")
	messages, channelHistoryErr := getChannelHistory(api, channel, startDate, endDate)

	if channelHistoryErr != nil {
		log.Printf("error getting channel history: %s\n", channelHistoryErr)
	}

	memoizedUserInfo := createUserMemoizer(api)
	for _, msg := range messages {

		user, err := memoizedUserInfo(msg.User)

		if err == nil {
			timestamp := strings.SplitN(msg.Timestamp, ".", -1)
			unixTimestamp, err := strconv.ParseInt(timestamp[0], 10, 64)
			if err != nil {
				panic(err)
			}
			time := time.Unix(unixTimestamp, 0)
			log.Printf("Fetching info for message in date: %s\n", time)
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
