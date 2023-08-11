package main

import (
	"fmt"
	"log"

	ctfd "github.com/pandatix/go-ctfd/api"
)

func main() {
	url := "http://127.0.0.1:8080"
	apiKey := "fec7c3341ed4624062f8b57dbebadf1bdd79a1b37f32420c281884c9cab855af"

	client := ctfd.NewClient(url, "", "", apiKey)

	files, err := client.GetChallengeFiles("1")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		content, err := client.GetFileContent(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[ ] File %d %s\n%s\n", file.ID, file.Location, content)
	}
}
