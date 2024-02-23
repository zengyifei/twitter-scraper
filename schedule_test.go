package twitterscraper_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFetchScheduledTweets(t *testing.T) {
	scheduled, err := testScraper.FetchScheduledTweets()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.Marshal(scheduled)
	fmt.Println(string(b))
}

func TestDeleteScheduledTweets(t *testing.T) {
	if err := testScraper.DeleteScheduledTweet("1760827700724355072"); err != nil {
		t.Error(err)
	}
}
