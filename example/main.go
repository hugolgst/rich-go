package main

import (
	"fmt"
	"time"

	"github.com/heroslender/rich-go/client"
)

func main() {

	err := client.Login("YOUR_DISCORD_APP_ID")
	if err != nil {
		panic(err)
	}

	err = client.SetActivity(client.Activity{
		State:      "hey",
		Details:    "i'm running on go",
		LargeImage: "Unknown",   // TODO: Add image
		LargeText:  "None",      // TODO: Add image alt
		SmallImage: "Unknown",   // TODO: Add image
		SmallText:  "NoneSmall", // TODO: Add image alt
		Party: &client.Party{
			ID:         "-1",
			Players:    15,
			MaxPlayers: 24,
		},
		Timestamps: &client.Timestamps{
			Start: time.Now().UnixNano() / 1e6,
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
