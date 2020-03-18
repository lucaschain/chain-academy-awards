package winning_categories

import (
	"sort"

	"github.com/lucaschain/chain-academy-awards/slack"
)

// MostBlockchains determines a winner based
// on the number of Blockchains of the message
// it always returns a slice because the first position may be tied
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
