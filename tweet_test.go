package twitterscraper_test

import (
	"testing"

	twitterscraper "github.com/imperatrona/twitter-scraper"
)

var testDeleteTweetId string

func TestCreateTweet(t *testing.T) {
	if skipAuthTest {
		t.Skip("Skipping test due to environment variable")
	}

	var err error
	var tweet *twitterscraper.Tweet
	tweet, err = testScraper.CreateTweet(twitterscraper.NewTweet{
		Text:   "i love hollywood 🖤",
		Medias: nil,
	})

	if tweet != nil {
		testDeleteTweetId = tweet.ID
	}
	if err != nil {
		t.Error(err)
	}
}

func TestCreateTweetWithMedia(t *testing.T) {
	if skipAuthTest {
		t.Skip("Skipping test due to environment variable")
	}

	var err error

	var video *twitterscraper.Media
	video, err = testScraper.UploadMedia("./photo.jpg")
	if err != nil {
		t.Error(err)
	}

	var photo *twitterscraper.Media
	photo, err = testScraper.UploadMedia("./video.mp4")
	if err != nil {
		t.Error(err)
	}

	var tweet *twitterscraper.Tweet
	tweet, err = testScraper.CreateTweet(twitterscraper.NewTweet{
		Text: "3 more seconds till i get 🖤",
		Medias: []*twitterscraper.Media{
			photo,
			video,
		},
	})

	if tweet != nil {
		testDeleteTweetId = tweet.ID
	}
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTweet(t *testing.T) {
	if skipAuthTest {
		t.Skip("Skipping test due to environment variable")
	}
	if testDeleteTweetId == "" {
		t.Skip("run TestCreateTweet before")
	}

	if err := testScraper.DeleteTweet(testDeleteTweetId); err != nil {
		t.Error(err)
	}
}

func TestCreateRetweet(t *testing.T) {
	if skipAuthTest {
		t.Skip("Skipping test due to environment variable")
	}
	if _, err := testScraper.CreateRetweet("1792634158977568997"); err != nil {
		t.Error(err)
	}
}

func TestDeleteRetweet(t *testing.T) {
	if skipAuthTest {
		t.Skip("Skipping test due to environment variable")
	}
	if err := testScraper.DeleteRetweet("1792634158977568997"); err != nil {
		t.Error(err)
	}
}
