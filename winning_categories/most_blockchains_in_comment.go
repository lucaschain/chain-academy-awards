package winning_categories

import (
	"sort"

	"github.com/lucaschain/chain-academy-awards/slack"
)

func maximumReplyBlockchains(message slack.Message) int {
	replies := message.Replies

	if len(replies) == 0 {
		return 0
	}

	// this can be done in a cheaper way if needed
	sort.SliceStable(replies, func(j, k int) bool {
		return replies[j].Blockchains > replies[k].Blockchains
	})

	return replies[0].Blockchains
}

// MostBlockchains determines a winner based
// on the number of Blockchains of the message
// it always returns a slice because the first position may be tied
func MostBlockchainsInComment(messages []slack.Message) []slack.Message {

	sort.SliceStable(messages, func(j, k int) bool {
		return maximumReplyBlockchains(messages[j]) > maximumReplyBlockchains(messages[k])
	})

	winnerScore := maximumReplyBlockchains(messages[0])

	winners := []slack.Message{}

	for _, msg := range messages {
		if maximumReplyBlockchains(msg) == winnerScore {
			winners = append(winners, msg)
		} else {
			break
		}
	}

	return winners
}
