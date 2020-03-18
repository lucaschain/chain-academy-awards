package winning_categories

import (
	"sort"

	"github.com/lucaschain/chain-academy-awards/slack"
)

func MostBlockchains(messages []slack.Message) []slack.Message {

	sort.SliceStable(messages, func(j, k int) bool {
		return messages[j].Blockchains > messages[k].Blockchains
	})

	winnerScore := messages[0].Blockchains

	winners := []slack.Message{}

	for _, msg := range messages {
		if msg.Blockchains == winnerScore {
			winners = append(winners, msg)
		} else {
			break
		}
	}

	return winners
}
