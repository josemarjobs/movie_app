package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Credit struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

func (c Credit) NameOrTitle() string {
	if c.Name != "" {
		return c.Name
	}
	return c.Title
}

type CreditSearchResults struct {
	Cast []Credit `json:"cast"`
}

func FetchCredits(actor *Actor) error {
	u := fmt.Sprintf("%s/person/%d/combined_credits?api_key=%s", ApiRoot, actor.ID, ApiKey)
	results := CreditSearchResults{}

	res, err := http.Get(u)
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&results)
	if err != nil {
		return err
	}

	actor.Credits = results.Cast
	return nil
}

func FilterCredits(actors []Actor) []Credit {
	credits := []Credit{}

	a := actors[0]
	al := len(actors)

	var wg sync.WaitGroup
	var m sync.Mutex
	for _, c := range a.Credits {
		wg.Add(1)
		go func(c Credit) {
			defer wg.Done()
			count := 1
			for _, actor := range actors[1:] {
				for _, actorCredit := range actor.Credits {
					if actorCredit.ID == c.ID {
						count += 1
						break
					}
				}
			}
			if count == al {
				m.Lock()
				credits = append(credits, c)
				m.Unlock()
			}

		}(c)
	}
	wg.Wait()
	return credits
}
