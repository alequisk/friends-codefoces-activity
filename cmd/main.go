package main

import (
	"net/http"
	"time"

	"github.com/alequisk/cf-friends-activity/internal/domain/checker"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	period_for_check := 5 * time.Minute
	c := checker.NewChecker(client, period_for_check)
	c.Run()
}
