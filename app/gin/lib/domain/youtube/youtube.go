package youtube

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	log "game-information/lib/domain/log"
)

// Video 動画情報
type Video struct {
	ID          string
	ChannelID   string
	Description string
	PublishedAt time.Time
	Title       string
}

// InterfaceYoutube インターフェース
type InterfaceYoutube interface {
	getService() *youtube.Service
	GetVideos(channelID string, query string, publishedAfter string, token string) *[]Video
}

// Youtube インスタンス
type Youtube struct {
	DeveloperKey string
}

func GetInstance() *Youtube {
	return &Youtube{DeveloperKey: os.Getenv("YOUTUBE_API_KEY")}
}

// getService serviceの生成
func (domain *Youtube) getService() *youtube.Service {
	client := &http.Client{
		Transport: &transport.APIKey{Key: domain.DeveloperKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Error(fmt.Sprintf("Error creating new YouTube client: %v", err), nil)
		return nil
	}
	return service
}

// GetVideos youtubeからデータ取得
func (domain *Youtube) GetVideos(channelID string, query string, publishedAfter string, token string) *[]Video {
	service := domain.getService()
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
		log.Error(fmt.Sprintf("Error making API call to search: %v", err.Error()), nil)
		return &[]Video{}
	}
	videos := []Video{}

	if response.NextPageToken != "" {
		videos = *domain.GetVideos(channelID, query, publishedAfter, response.NextPageToken)
	}

	for _, item := range response.Items {
		publishedAt, _ := time.Parse("2006-01-02T15:04:05Z", item.Snippet.PublishedAt)
		videos = append(videos, Video{item.Id.VideoId, item.Snippet.ChannelId, item.Snippet.Description, publishedAt, item.Snippet.Title})
	}
	return &videos
}
