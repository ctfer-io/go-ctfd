package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ctfer-io/go-ctfd/api"
)

const (
	url    = "http://localhost:8000"
	apiKey = "ctfd_xxx"
)

func main() {
	ctx := context.Background()
	cli := api.NewClient(url, "", "", apiKey)

	chall, err := PostChallenges(cli, &PluginPostChallengesParams{
		PostChallengesParams: api.PostChallengesParams{
			Name:        "My challenge",
			Description: "Super duper description",
			Category:    "Misc",
			Value:       500,
			Type:        "standard",
		},
		NewField1: "some content",
		NewField2: map[string]bool{
			"one": true,
			"two": false,
		},
	}, api.WithContext(ctx))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[%d] new_field3: %v\n", chall.ID, chall.NewField3)
}

type PluginPostChallengesParams struct {
	api.PostChallengesParams

	NewField1 string          `json:"new_field1"`
	NewField2 map[string]bool `json:"new_field2,omitempty"`
}

// ... other structs required for your use case

type MyPluginChallenge struct {
	api.Challenge

	NewField1 string          `json:"new_field1"`
	NewField2 map[string]bool `json:"new_field2,omitempty"`
	NewField3 int             `json:"new_field3"`
}

func PostChallenges(client *api.Client, params *PluginPostChallengesParams, opts ...api.Option) (*MyPluginChallenge, error) {
	if params == nil {
		params = &PluginPostChallengesParams{}
	}
	// ... default any value if required, check values (e.g. integer boundaries)

	chall := &MyPluginChallenge{}
	if err := client.Post("/plugins/my_plugin/challenges", params, chall, opts...); err != nil {
		return nil, err
	}
	return chall, nil
}

// ... other methods required for your use case

// you can even go further and implement your custom client that
// fit your plugin API needs ! :confetti_ball:
