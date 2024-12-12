package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ctfer-io/go-ctfd/api"
)

const (
	url    = "http://localhost:8000"
	apiKey = "ctfd_xxx"
)

func main() {
	ctx := context.Background()
	cli := api.NewClient(url, "", "", apiKey)

	chall, err := PostChallenges(cli, ctx, &MyPluginPostChallengeParams{
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
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[%d] new_field3: %v\n", chall.ID, chall.NewField3)
}

type MyPluginPostChallengeParams struct {
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

func PostChallenges(client *api.Client, ctx context.Context, params *MyPluginPostChallengeParams) (*MyPluginChallenge, error) {
	if params == nil {
		params = &MyPluginPostChallengeParams{}
	}
	// ... default any value if required, check values (e.g. integer boundaries)

	// Create your API object
	chall := &MyPluginChallenge{}

	// Make the call
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(params); err != nil {
		return nil, err
	}
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/plugins/my_plugin/challenges", nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response
	resp := api.Response{
		Data: chall,
	}
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	// ... do any API check you would like to

	return chall, nil
}

// ... other methods required for your use case

// you can even go further and implement your custom client that
// fit your plugin API needs ! :confetti_ball:
