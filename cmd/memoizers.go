package cmd

import (
	"github.com/slack-go/slack"
)

func createUserMemoizer(api *slack.Client) func(string) (*slack.User, error) {
	cache := make(map[string]*slack.User)

	return func(userID string) (*slack.User, error) {
		if cache[userID] == nil {
			user, err := api.GetUserInfo(userID)

			if err != nil {
				return nil, err
			}

			cache[userID] = user
		}

		return cache[userID], nil
	}
}
