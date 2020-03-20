package winning_categories

import (
	"github.com/lucaschain/chain-academy-awards/slack"
)

type ScoreTable = map[string]int

func findWinnerByScore(scoreTable ScoreTable) string {
	maximumScore := 0
	winner := ""

	for user, score := range scoreTable {
		if score > maximumScore {
			maximumScore = score
			winner = user
		}
	}

	return winner
}

func countScores(messages []slack.Message) ScoreTable {
	scoreTable := ScoreTable{}

	for _, msg := range messages {
		userScore, ok := scoreTable[msg.User]
		if ok {
			scoreTable[msg.User] = userScore + 1
			continue
		}

		scoreTable[msg.User] = 1
	}

	return scoreTable
}

func winnerMessages(winner string, messages []slack.Message) []slack.Message {
	winnerMessages := []slack.Message{}

	for _, msg := range messages {
		if msg.User == winner {
			winnerMessages = append(winnerMessages, msg)
		}
	}

	return winnerMessages
}

func MostActiveUser(messages []slack.Message) []slack.Message {
	scoreTable := countScores(messages)

	winner := findWinnerByScore(scoreTable)

	return winnerMessages(winner, messages)
}
