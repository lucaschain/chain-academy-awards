package winning_categories

import (
	"sort"

	"github.com/lucaschain/chain-academy-awards/slack"
)

// MostReplies determines a winner based
// on the number of Replies in the message
// it always returns a slice because the first position may be tied
func MostReplies(messages []slack.Message) []slack.Message {

	sort.SliceStable(messages, func(j, k int) bool {
		return len(messages[j].Replies) > len(messages[k].Replies)
	})

	winnerScore := len(messages[0].Replies)

	winners := []slack.Message{}

	for _, msg := range messages {
		if len(msg.Replies) == winnerScore {
			winners = append(winners, msg)
		} else {
			break
		}
	}

	return winners
}
