package slack

import (
	"log"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

func findAndJoinChannel(api *slack.Client, channelName string) *slack.Channel {
	conversationsParam := &slack.GetConversationsParameters{
		Limit: 1000, // TODO fix this
		Types: []string{"public_channel"},
	}
	channels, _, err := api.GetConversations(conversationsParam)

	if err != nil {
		log.Printf("%s\n", err)
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
		log.Printf("error joining: %s\n", err)
		return nil
	}

	return &channel
}

func getChannelHistory(api *slack.Client, channel *slack.Channel, from time.Time, to time.Time) ([]slack.Message, error) {

	messages := []slack.Message{}
	cursor := ""

	for {
		log.Println("new history batch")
		conversationHistoryParams := &slack.GetConversationHistoryParameters{
			ChannelID: channel.ID,
			Cursor:    cursor,
			Oldest:    strconv.FormatInt(from.Unix(), 10),
			Latest:    strconv.FormatInt(to.Unix(), 10),
		}

		history, channelHistoryErr := api.GetConversationHistory(conversationHistoryParams)
		if channelHistoryErr != nil {
			log.Printf("error getting channel history: %s\n", channelHistoryErr)
			continue
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
