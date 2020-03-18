package cmd

import (
	"fmt"

	"github.com/slack-go/slack"
)

func repliesForMessage(api *slack.Client, channel *slack.Channel, msg slack.Message, userMemoizer func(string) (*slack.User, error)) []Reply {
	replies := []Reply{}

	for _, reply := range msg.Replies {
		ts := reply.Timestamp
		conversationRepliesParams := &slack.GetConversationRepliesParameters{
			ChannelID: channel.ID,
			Timestamp: ts,
		}
		replyMessages, _, _, err := api.GetConversationReplies(conversationRepliesParams)

		if err == nil {
			for _, replyMsg := range replyMessages {
				userInfo, _ := userMemoizer(replyMsg.User)

				replies = append(replies, NewReply(
					replyMsg.Text,
					blockchainsForMessage(api, channel, replyMsg.Timestamp),
					userInfo.Name,
				))
			}
		} else {
			fmt.Printf("error in thread: %s\n", err)
		}
	}

	return replies
}
