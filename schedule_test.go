package twitterscraper_test

import (
	"testing"
)

func TestFetchScheduledTweets(t *testing.T) {
	_, err := testScraper.FetchScheduledTweets()
	if err != nil {
		t.Error(err)
	}
}
