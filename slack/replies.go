package slack

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
)

func repliesForMessage(api *slack.Client, channel *slack.Channel, msg slack.Message, userMemoizer func(string) (*slack.User, error)) []Reply {
	replies := []Reply{}

	log.Println("fetching message replies")
	for _, reply := range msg.Replies {
		ts := reply.Timestamp
		conversationRepliesParams := &slack.GetConversationRepliesParameters{
			ChannelID: channel.ID,
			Timestamp: ts,
		}
		replyMessages, _, _, err := api.GetConversationReplies(conversationRepliesParams)

		if err == nil {
			for _, replyMsg := range replyMessages {
				log.Println("fetching reply")
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
