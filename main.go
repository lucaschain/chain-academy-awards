package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/lucaschain/chain-academy-awards/slack"
	"github.com/lucaschain/chain-academy-awards/winning_categories"
)

type AwardCounter = func([]slack.Message) []slack.Message

func showAwardResult(awardName string, messages []slack.Message, counter AwardCounter) {
	fmt.Printf("And the winners for %s are:\n", awardName)

	winners := counter(messages)

	for _, w := range winners {
		w.Print()
	}
}

func generateOutput() {
	messages := []slack.Message{}

	slack.Fetch("development", func(message slack.Message) {
		messages = append(messages, message)
	})

	json, _ := json.Marshal(messages)
	_ = ioutil.WriteFile("output.json", json, 0644)
}

func main() {
	file, err := ioutil.ReadFile("output.json")

	if err == nil {
		messages := []slack.Message{}
		_ = json.Unmarshal([]byte(file), &messages)

		showAwardResult("most blockchains", messages, winning_categories.MostBlockchains)
	} else {
		generateOutput()
	}
}
