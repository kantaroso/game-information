package youtube

import (
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

// Movie 動画情報
type Movie struct {
	ID          string
	ChannelID   string
	Description string
	PublishedAt time.Time
	Title       string
}

// getService serviceの生成
func getService() *youtube.Service {
	developerKey := os.Getenv("YOUTUBE_API_KEY")
	log.Print(developerKey)
	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}
	return service
}

// GetMovieList youtubeからデータ取得
func GetMovieList() *[]Movie {

	service := getService()
	// test
	channelID := "UCVj6bacaTQ6r6Dsv0frD_sA"
	order := "date"
	query := "PV OR ティザームービー"
	var maxResult int64 = 50

	part := []string{"id", "snippet"}
	call := service.Search.List(part).ChannelId(channelID).Order(order).Q(query).MaxResults(maxResult)

	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call to search: %v", err.Error())
	}
	movies := []Movie{}
	for _, item := range response.Items {
		publishedAt, _ := time.Parse("2006-01-02T15:04:05Z", item.Snippet.PublishedAt)
		movies = append(movies, Movie{item.Id.VideoId, item.Snippet.ChannelId, item.Snippet.Description, publishedAt, item.Snippet.Title})
	}
	return &movies
}
