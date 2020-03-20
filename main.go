package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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

	slack.Fetch("chain-academy", func(message slack.Message) {
		messages = append(messages, message)
	})

	json, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("could not marshal json:", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("output.json", json, 0644)

	if err != nil {
		fmt.Println("could not save file:", err)
		os.Exit(1)
	}
}

func showAwards() {
	file, err := ioutil.ReadFile("output.json")

	if err != nil {
		fmt.Println("could not load file:", err)
		os.Exit(1)
	}

	messages := []slack.Message{}
	err = json.Unmarshal([]byte(file), &messages)
	if err != nil {
		fmt.Println("could not unmarshal json:", err)
		os.Exit(1)
	}

	awards := map[string]AwardCounter{
		"most blockchains":            winning_categories.MostBlockchains,
		"most replies":                winning_categories.MostReplies,
		"most blockchains in thread":  winning_categories.MostBlockchainsInThread,
		"most blockchains in comment": winning_categories.MostBlockchainsInComment,
	}

	for name, counter := range awards {
		showAwardResult(name, messages, counter)
	}
}

func main() {
	_, err := os.Stat("output.json")

	if err != nil {
		generateOutput()
	}

	showAwards()
}
