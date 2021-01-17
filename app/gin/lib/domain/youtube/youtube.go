package youtube

import (
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

// Video 動画情報
type Video struct {
	ID          string
	ChannelID   string
	Description string
	PublishedAt time.Time
	Title       string
}

// getService serviceの生成
func getService() *youtube.Service {
	developerKey := os.Getenv("YOUTUBE_API_KEY")
	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}
	return service
}

// GetVideos youtubeからデータ取得
func GetVideos(channelID string, query string, publishedAfter string, token string) *[]Video {
	service := getService()
	// test
	order := "date"
	var maxResult int64 = 50
	targettype := "video"

	if publishedAfter == "" {
		// 最古まで指定したいのでyoutubeのサービス開始日を利用
		tmpPublishedAfter, _ := time.Parse("2006/01/02", "2005/12/15")
		publishedAfter = tmpPublishedAfter.Format(time.RFC3339)
	}

	part := []string{"id", "snippet"}
	call := service.Search.List(part).ChannelId(channelID).Type(targettype).Order(order).Q(query).MaxResults(maxResult).PublishedAfter(publishedAfter)
	if token != "" {
		call = service.Search.List(part).ChannelId(channelID).Type(targettype).Order(order).Q(query).MaxResults(maxResult).PublishedAfter(publishedAfter).PageToken(token)
	}

	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call to search: %v", err.Error())
	}
	videos := []Video{}

	if response.NextPageToken != "" {
		videos = *GetVideos(channelID, query, publishedAfter, response.NextPageToken)
	}

	for _, item := range response.Items {
		publishedAt, _ := time.Parse("2006-01-02T15:04:05Z", item.Snippet.PublishedAt)
		videos = append(videos, Video{item.Id.VideoId, item.Snippet.ChannelId, item.Snippet.Description, publishedAt, item.Snippet.Title})
	}
	return &videos
}
