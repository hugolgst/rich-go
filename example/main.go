package main

import (
	"fmt"
	"os"
	"time"

	"github.com/donovansolms/rich-go-redo/client"
)

func main() {
	fmt.Println("Test presence")

	err := client.Login("530821687864983554")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = client.SetActivity(client.Activity{
		State:   "hey",
		Details: "i'm running on go",
		Assets: client.Assets{
			LargeImage: "Unknown",   // TODO: Add image
			LargeText:  "None",      // TODO: Add image alt
			SmallImage: "Unkown",    // TODO: Add image
			SmallText:  "NoneSmall", // TODO: Add image alt
		},
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	time.Sleep(time.Second * 5)
}
