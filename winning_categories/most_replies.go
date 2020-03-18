package winning_categories

import (
	"sort"

	"github.com/lucaschain/chain-academy-awards/slack"
)

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
