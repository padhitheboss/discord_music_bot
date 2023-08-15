package model

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/padhitheboss/sangeet/pkg/config"
	"google.golang.org/api/youtube/v3"
)

type Song struct {
	Id       int
	Title    string
	Url      string
	AudioUrl string
	Duration string
	AddedOn  time.Time
	AddedBy  string
}
type PlayQueue struct {
	Lock        sync.Mutex
	CurId       int
	Playlist    []Song
	LastUpdated time.Time
	CreatedAt   time.Time
}

func searchYouTubeVideo(query string) (*youtube.Video, error) {
	searchResponse, err := config.YoutubeService.Search.List([]string{"snippet"}).Q(query).MaxResults(1).Type("video").Do()
	if err != nil {
		return &youtube.Video{}, err
	}

	if len(searchResponse.Items) > 0 {
		videoID := searchResponse.Items[0].Id.VideoId
		videoDetails, err := config.YoutubeService.Videos.List([]string{"snippet", "contentDetails"}).Id(videoID).Do()
		if err != nil {
			return &youtube.Video{}, err
		}
		return videoDetails.Items[0], nil
	}

	return &youtube.Video{}, fmt.Errorf("no result found")
}

func (s *Song) getSongDetails(title string) {
	videoDetails, err := searchYouTubeVideo(title)
	if err != nil {
		log.Println(err)
		return
	}
	s.Url = "https://www.youtube.com/watch?v=" + videoDetails.Id
	s.Title = videoDetails.Snippet.Title
	s.Duration = videoDetails.ContentDetails.Duration
	s.AddedOn = time.Now()

}
func (s *Song) CreateRequest(title string, addedBy string) {
	s.Title = title
	s.AddedBy = addedBy
	s.getSongDetails(s.Title)
}

func (p *PlayQueue) AddAtBack(s Song) {
	if p.CurId == 0 {
		p.CreatedAt = time.Now()
	}
	p.CurId++
	s.Id = p.CurId
	p.LastUpdated = time.Now()
	p.Playlist = append(p.Playlist, s)
}

func (p *PlayQueue) ClearQueue() {
	if len(p.Playlist) == 0 {
		log.Println("Queue is Already Empty")
		return
	}
	p.Playlist = p.Playlist[:0]
}

func (p *PlayQueue) AddAtFront(s Song) {
	p.Playlist = append([]Song{s}, p.Playlist...)
}
func (p *PlayQueue) RemoveFromFront() {
	if len(p.Playlist) == 0 {
		log.Println("Queue is Already Empty")
		return
	}
	p.Playlist = p.Playlist[1:]
}
