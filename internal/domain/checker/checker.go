package checker

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gen2brain/beeep"
)

const (
	API_URL = "https://codeforces.com/api/user.status?handle=%s&from=1&count=1"
)

type Checker struct {
	client      *http.Client
	periodCheck time.Duration
}

func (c Checker) fectchSubmissionsByHandles() []uint64 {
	var submissions_id = make([]uint64, len(handles))

	for index, handle := range handles {
		submission_id, err := c.retrieveLastSubmission(handle)
		if err != nil {
			submissions_id[index] = 0
		} else {
			submissions_id[index] = submission_id
		}

		time.Sleep(4 * time.Second)
	}
	return submissions_id
}

type submissionHTTPResponse struct {
	Status string       `json:"status"`
	Result []Submission `json:"result"`
}

func (c *Checker) retrieveLastSubmission(handle string) (uint64, error) {
	r, err := c.client.Get(fmt.Sprintf(API_URL, handle))

	if err != nil {
		return 0, err
	}

	defer r.Body.Close()
	sub := &submissionHTTPResponse{}

	if err = json.NewDecoder(r.Body).Decode(sub); err != nil {
		return 0, err
	}

	if sub.Status != "OK" {
		return 0, errors.New("request API failed")
	}

	return uint64(sub.Result[0].ID), nil
}

func (c *Checker) Run() {
	log.Println("Populate submissions...")
	subs := c.fectchSubmissionsByHandles()

	for {
		time.Sleep(c.periodCheck)
		log.Println("Updating submissions ID.")

		var friends_coding []string
		updated_subs := c.fectchSubmissionsByHandles()
		for i := range updated_subs {
			if subs[i] != updated_subs[i] {
				friends_coding = append(friends_coding, handles[i])
			}
		}

		subs = updated_subs

		if len(friends_coding) > 0 {
			log.Println("Notify that have new online friends")
			notify(friends_coding)
		}
	}
}

func notify(friends []string) {
	var names string
	for i, name := range friends {
		if i > 0 {
			if i+1 == len(friends) {
				names += " e "
			} else {
				names += ", "
			}
		}
		names += name
	}

	verb := "está"
	if len(friends) > 1 {
		verb = "estão"
	}
	message := fmt.Sprintf("%s %s condando neste momento,", names, verb)
	beeep.Notify("Amigo(s) on-line", message, "assets/warning.png")
}

func NewChecker(client *http.Client, period time.Duration) *Checker {
	return &Checker{
		client:      client,
		periodCheck: period,
	}
}
