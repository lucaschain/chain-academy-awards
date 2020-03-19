package slack

import (
	"log"

	"github.com/slack-go/slack"
)

func blockchainsForMessage(api *slack.Client, channel *slack.Channel, ts string) int {
	msgRef := slack.NewRefToMessage(channel.ID, ts)
	retries := 0
	for {
		msgReactions, err := api.GetReactions(msgRef, slack.NewGetReactionsParameters())

		if err != nil {
			log.Printf("error getting reactions: %s, retry %d\n", err, retries)
			retries += 1
			continue
		}

		blockchains := 0
		for _, r := range msgReactions {
			if r.Name == "blockchain" {
				blockchains = r.Count
			}
		}

		log.Println("done")

		return blockchains
	}
}
