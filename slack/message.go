package slack

import (
	"fmt"
)

type Awardable struct {
	Text        string
	Blockchains int
	User        string
}

type Message struct {
	Awardable
	Permalink string
	Replies   []Reply
}

type Reply struct {
	Awardable
}

func NewReply(text string, blockchains int, user string) Reply {
	reply := Reply{}

	reply.Text = text
	reply.Blockchains = blockchains
	reply.User = user

	return reply
}

func NewMessage(text string, blockchains int, user string, replies []Reply, permalink string) Message {
	msg := Message{}

	msg.Text = text
	msg.Blockchains = blockchains
	msg.User = user
	msg.Replies = replies
	msg.Permalink = permalink

	return msg
}

func (message *Message) Print() {
	fmt.Println(message.Permalink)
	fmt.Printf("%s: %s (%d blockchains)\n", message.User, message.Text, message.Blockchains)
	for _, reply := range message.Replies {
		fmt.Printf("    - %s: %s (%d blockchains)\n", reply.User, reply.Text, reply.Blockchains)
	}
}
