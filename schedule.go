package twitterscraper

import (
	"net/url"
	"strings"
	"time"
)

type scheduleTweet struct {
	RestID         string `json:"rest_id"`
	SchedulingInfo struct {
		ExecuteAt int64  `json:"execute_at"`
		State     string `json:"state"`
	} `json:"scheduling_info"`
	TweetCreateRequest struct {
		Type                      string        `json:"type"`
		Status                    string        `json:"status"`
		ExcludeReplyUserIds       []interface{} `json:"exclude_reply_user_ids"`
		MediaIds                  []interface{} `json:"media_ids"`
		AutoPopulateReplyMetadata bool          `json:"auto_populate_reply_metadata"`
	} `json:"tweet_create_request"`
	MediaEntities []struct {
		MediaKey  string `json:"media_key"`
		MediaInfo struct {
			Typename          string `json:"__typename"`
			OriginalImgURL    string `json:"original_img_url"`
			OriginalImgWidth  int    `json:"original_img_width"`
			OriginalImgHeight int    `json:"original_img_height"`
			DurationMillis    int    `json:"duration_millis"`
			Variants          []struct {
				ContentType string `json:"content_type"`
				Bitrate     int    `json:"bit_rate,omitempty"`
				URL         string `json:"url"`
			} `json:"variants"`
			AspectRatio struct {
				Numerator   int `json:"numerator"`
				Denominator int `json:"denominator"`
			} `json:"aspect_ratio"`
			PreviewImage struct {
				OriginalImgURL    string `json:"original_img_url"`
				OriginalImgWidth  int    `json:"original_img_width"`
				OriginalImgHeight int    `json:"original_img_height"`
			} `json:"preview_image"`
		} `json:"media_info"`
	} `json:"media_entities,omitempty"`
}

func (result *scheduleTweet) parse() *ScheduledTweet {
	tweet := &ScheduledTweet{
		ID:        result.RestID,
		State:     result.SchedulingInfo.State,
		ExecuteAt: time.Unix(result.SchedulingInfo.ExecuteAt/1000, 0),
		Text:      result.TweetCreateRequest.Status,
	}

	for _, media := range result.MediaEntities {
		k := strings.Split(media.MediaKey, "_")
		key := k[len(k)-1]

		if media.MediaInfo.Typename == "ApiVideo" {
			video := Video{
				ID:      key,
				Preview: media.MediaInfo.PreviewImage.OriginalImgURL,
			}

			maxBitrate := 0
			for _, variant := range media.MediaInfo.Variants {
				if variant.Bitrate > maxBitrate {
					video.URL = strings.TrimSuffix(variant.URL, "?tag=10")
					maxBitrate = variant.Bitrate
				}
			}

			tweet.Videos = append(tweet.Videos, video)
		} else if media.MediaInfo.Typename == "ApiGif" {
			gif := GIF{
				ID:      key,
				Preview: media.MediaInfo.PreviewImage.OriginalImgURL,
			}

			maxBitrate := 0
			for _, variant := range media.MediaInfo.Variants {
				if variant.Bitrate >= maxBitrate {
					gif.URL = variant.URL
					maxBitrate = variant.Bitrate
				}
			}
			tweet.GIFs = append(tweet.GIFs, gif)
		} else if media.MediaInfo.Typename == "ApiImage" {
			tweet.Photos = append(tweet.Photos, Photo{
				ID:  key,
				URL: media.MediaInfo.OriginalImgURL,
			})
		}
	}

	return tweet
}

type scheduleTweets struct {
	Data struct {
		Viewer struct {
			ScheduledTweetList []scheduleTweet `json:"scheduled_tweet_list"`
		} `json:"viewer"`
	} `json:"data"`
}

func (timeline *scheduleTweets) parseTweets() []*ScheduledTweet {
	var tweets []*ScheduledTweet

	for _, entry := range timeline.Data.Viewer.ScheduledTweetList {
		if tweet := entry.parse(); tweet != nil {
			tweets = append(tweets, tweet)
		}
	}
	return tweets
}

// FetchScheduledTweets gets scheduled tweets via the Twitter frontend GraphQL API.
func (s *Scraper) FetchScheduledTweets() ([]*ScheduledTweet, error) {
	req, err := s.newRequest("GET", "https://twitter.com/i/api/graphql/ITtjAzvlZni2wWXwf295Qg/FetchScheduledTweets")
	if err != nil {
		return nil, err
	}

	variables := map[string]interface{}{
		"ascending": true,
	}

	query := url.Values{}
	query.Set("variables", mapToJSONString(variables))
	req.URL.RawQuery = query.Encode()

	var timeline scheduleTweets
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	tweets := timeline.parseTweets()
	return tweets, nil
}
