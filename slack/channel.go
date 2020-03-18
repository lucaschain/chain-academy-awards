package slack

import (
	"fmt"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

func findAndJoinChannel(api *slack.Client, channelName string) *slack.Channel {
	conversationsParam := &slack.GetConversationsParameters{
		Limit: 10,
	}
	channels, _, err := api.GetConversations(conversationsParam)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	var channel slack.Channel
	for _, c := range channels {

		if c.Name == channelName {
			channel = c
			break
		}
	}

	if channel.ID == "" {
		return nil
	}

	_, _, _, err = api.JoinConversation(channel.ID)

	if err != nil {
		fmt.Printf("error joining: %s\n", err)
		return nil
	}

	return &channel
}

func getChannelHistory(api *slack.Client, channel *slack.Channel, from time.Time, to time.Time) ([]slack.Message, error) {

	messages := []slack.Message{}
	cursor := ""

	for {
		conversationHistoryParams := &slack.GetConversationHistoryParameters{
			ChannelID: channel.ID,
			Cursor:    cursor,
			Oldest:    strconv.FormatInt(from.Unix(), 10),
			Latest:    strconv.FormatInt(to.Unix(), 10),
		}

		history, channelHistoryErr := api.GetConversationHistory(conversationHistoryParams)
		if channelHistoryErr != nil {
			return messages, fmt.Errorf("error getting channel history: %s\n", channelHistoryErr)
		}

		for _, msg := range history.Messages {
			if msg.SubType == "" {
				messages = append(messages, msg)
			}
		}

		cursor = history.ResponseMetaData.NextCursor

		if history.HasMore == false {
			break
		}
	}

	return messages, nil
}
