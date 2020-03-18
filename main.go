package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/lucaschain/chain-academy-awards/cmd"
)

func main() {
	messages := []cmd.Message{}

	cmd.Fetch("development", func(message cmd.Message) {
		messages = append(messages, message)

		fmt.Printf("%s: %s (%d blockchains)\n", message.User, message.Text, message.Blockchains)
		for _, reply := range message.Replies {
			fmt.Printf("    - %s: %s (%d blockchains)\n", reply.User, reply.Text, reply.Blockchains)
		}
	})

	json, _ := json.Marshal(messages)
	_ = ioutil.WriteFile("output.json", json, 0644)
}
