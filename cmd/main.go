package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/alequisk/cf-friends-activity/internal/domain/checker"
)

func main() {
	handles := os.Args[1:]

	if len(handles) == 0 {
		fmt.Println("You must specify the handles that will be monitored after execute the executable")
		fmt.Println("Example: .\\main.exe handle1 handle2 (windows)  or   ./main handle1 handle2 (linux)")
		os.Exit(1)
	}

	fmt.Print("Fetching submissions of handles: ")
	fmt.Println(strings.Join(handles, ","))

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	period_for_check := 5 * time.Minute
	c := checker.NewChecker(client, period_for_check, handles)
	c.Run()
}
