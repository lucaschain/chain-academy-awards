package cmd

import (
	"fmt"

	"github.com/slack-go/slack"
)

func blockchainsForMessage(api *slack.Client, channel *slack.Channel, ts string) int {
	msgRef := slack.NewRefToMessage(channel.ID, ts)
	msgReactions, err := api.GetReactions(msgRef, slack.NewGetReactionsParameters())

	if err != nil {
		fmt.Printf("Error getting reactions: %s\n", err)
		return 0
	}

	blockchains := 0
	for _, r := range msgReactions {
		if r.Name == "sunglasses" {
			blockchains = r.Count
		}
	}
	return blockchains
}
