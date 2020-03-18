package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lucaschain/chain-academy-awards/slack"
)

func generateOutput() {
	messages := []slack.Message{}

	slack.Fetch("development", func(message slack.Message) {
		messages = append(messages, message)

		fmt.Printf("%s: %s (%d blockchains)\n", message.User, message.Text, message.Blockchains)
		for _, reply := range message.Replies {
			fmt.Printf("    - %s: %s (%d blockchains)\n", reply.User, reply.Text, reply.Blockchains)
		}
	})

	json, _ := json.Marshal(messages)
	_ = ioutil.WriteFile("output.json", json, 0644)
}

func main() {
	if _, err := os.Stat("output.json"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		generateOutput()
	}
}
