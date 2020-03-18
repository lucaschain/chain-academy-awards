package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/slack-go/slack"
)

func Fetch(channelName string, resultCallback func(Message)) {
	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true))

	channel := findAndJoinChannel(api, channelName)

	if channel == nil {
		fmt.Println("Channel not found :'(")
		return
	}

	startDate := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	messages, channelHistoryErr := getChannelHistory(api, channel, startDate, endDate)

	if channelHistoryErr != nil {
		fmt.Printf("error getting channel history: %s\n", channelHistoryErr)
	}

	memoizedUserInfo := createUserMemoizer(api)
	for _, msg := range messages {

		user, err := memoizedUserInfo(msg.User)

		if err == nil {
			resultCallback(NewMessage(
				msg.Text,
				blockchainsForMessage(api, channel, msg.Timestamp),
				user.Name,
				repliesForMessage(api, channel, msg, memoizedUserInfo),
			))
		}

	}
}
