package main

import (
	"fmt"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/hugolgst/rich-go/client"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	procEnumWindows  = user32.NewProc("EnumWindows")
	procGetWindowText = user32.NewProc("GetWindowTextW")
)

func getTidalTitle() string {
	var title string
	cb := syscall.NewCallback(func(hwnd uintptr, lparam uintptr) uintptr {
		b := make([]uint16, 200)
		_, _, _ = procGetWindowText.Call(hwnd, uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)))
		t := syscall.UTF16ToString(b)

		if strings.Contains(t, " - ") {
			// EXTENDED IGNORE LIST: Added your specific browser and Gemini
			ignore := []string{
				"Google Chrome", "Edge", "Settings", "File Explorer", 
				"NVIDIA", "Discord", "Command Prompt", "main.go", "rich-go",
				"Zen Browser", "Google Gemini", "Terminal",
			}
			
			isSystem := false
			lowerTitle := strings.ToLower(t)
			for _, word := range ignore {
				if strings.Contains(lowerTitle, strings.ToLower(word)) {
					isSystem = true
					break
				}
			}

			if !isSystem {
				title = t
				return 0 
			}
		}
		return 1
	})
	_, _, _ = procEnumWindows.Call(cb, 0)
	return title
}

func main() {
	err := client.Login("1495530318357205094")
	if err != nil {
		fmt.Println("Discord Login Failed:", err)
		return
	}

	fmt.Println("Tidal Auto-Presence is running!")

	lastTitle := ""
	for {
		fullTitle := getTidalTitle()

		if fullTitle != "" && fullTitle != lastTitle {
			fmt.Println("Now Playing:", fullTitle)
			lastTitle = fullTitle

			parts := strings.Split(fullTitle, " - ")
			song := parts[0]
			artist := "Various Artists"
			if len(parts) > 1 {
				artist = parts[1]
			}

			err = client.SetActivity(client.Activity{
				Details:    song,
				State:      "by " + artist,
				LargeImage: "tidal_logo", 
				LargeText:  "Listening on Tidal",
			})
		} else if fullTitle == "" {
			fmt.Println("Looking for music...")
		}

		time.Sleep(time.Second * 5) 
	}
}