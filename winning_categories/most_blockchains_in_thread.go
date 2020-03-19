package winning_categories

import (
	"sort"

	"github.com/lucaschain/chain-academy-awards/slack"
)

func countBlockchainsForThread(message slack.Message) int {
	count := message.Blockchains

	for _, reply := range message.Replies {
		count += reply.Blockchains
	}

	return count
}

// MostBlockchainsInThread determines a winner based
// on the number of Blockchains of the whole thread
// it always returns a slice because the first position may be tied
func MostBlockchainsInThread(messages []slack.Message) []slack.Message {

	sort.SliceStable(messages, func(j, k int) bool {
		return countBlockchainsForThread(messages[j]) > countBlockchainsForThread(messages[k])
	})

	winnerScore := countBlockchainsForThread(messages[0])

	winners := []slack.Message{}

	for _, msg := range messages {
		if countBlockchainsForThread(msg) == winnerScore {
			winners = append(winners, msg)
		} else {
			break
		}
	}

	return winners
}
