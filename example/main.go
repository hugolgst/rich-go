package main

import (
	"fmt"
	"github.com/ananagame/rich-go/client"
	"time"
)

func main() {
	err := client.Login("DISCORD_APP_ID")
	if err != nil {
		panic(err)
	}

	err = client.SetActivity(client.Activity{
		State:      "Heyy!!!",
		Details:    "I'm running on rich-go :)",
		LargeImage: "largeimageid",
		LargeText:  "This is the large image :D",
		SmallImage: "smallimageid",
		SmallText:  "And this is the small image",
		Party: &client.Party{
			ID:         "-1",
			Players:    15,
			MaxPlayers: 24,
		},
		Timestamps: &client.Timestamps{
			Start: time.Now(),
		},
	})

	if err != nil {
		panic(err)
	}

	// Discord will only show the presence if the app is running
	// Sleep for a few seconds to see the update
	fmt.Println("Sleeping...")
	time.Sleep(time.Second * 10)
}
