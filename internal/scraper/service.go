package scraper

import (
	"fmt"

	"github.com/matmazurk/Kononobot/internal/rest"
	"github.com/matmazurk/Kononobot/pkg/persistance"
)

const (
	location = "Europe/Warsaw"
)

var (
	channels = map[string]string{
		"Mleczny człowiek": "UCPhnxJMplQz9A1DstzRDlEA",
		"Wojtek z bombasu": "UCa3y6lbZByfJ2tSPZ9XrjRw",
	}
)

type Config struct {
	ApiURL string `env:"API_URL"`
	ApiKey string `env:"API_KEY"`
}

type Service struct {
	config Config
	db     persistance.Repository
}

func (s Service) Run() {

}

func (s Service) insertAllFilmsForEachChannel() {
	for _, channelID := range channels {
		latestFilm := s.db.GetLatestFilm(channelID)
		remainingFilms, err := rest.GetAllFilmsAfter(s.config.ApiURL, s.config.ApiKey, channelID, latestFilm.PublishTime)
		if err != nil {
			fmt.Println(err)
		}

		for _, filmList := range remainingFilms {
			for _, film := range filmList.Items {
				persFilm := persistance.Film{
					ID:           film.Id.VideoId,
					PublishedAt:  film.Snippet.PublishedAt,
					ChannelID:    film.Snippet.ChannelId,
					Title:        film.Snippet.Title,
					Description:  film.Snippet.Description,
					ChannelTitle: film.Snippet.ChannelId,
					PublishTime:  film.Snippet.PublishTime,
				}
				s.db.InsertFilm(persFilm)
			}
		}
	}
}
