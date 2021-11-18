package services

import (
	"context"
	"time"

	"github.com/matmazurk/Kononobot/pkg/persistence"
)

type Handler interface{}
type Persistence interface {
	InsertChannel(
		ctx context.Context,
		id, title string,
		viewCount int64,
		subscriberCount, videoCount int32) error
	GetChannel(ctx context.Context, id string) (persistence.Channel, error)
	InsertFilm(
		ctx context.Context,
		publishedAt time.Time,
		channelId, title, description string,
		viewCount int64,
		likeCount, dislikeCount, commentCount int32) error
	GetFilmsForChannel(ctx context.Context, channelId string) ([]persistence.Film, error)
}

type kbotService struct {
	handler     Handler
	persistence Persistence
}

func NewKBot(h Handler, p Persistence) kbotService {
	return kbotService{
		handler:     h,
		persistence: p,
	}
}

func (s kbotService) Serve() error {
	s.persistence.InsertChannel(context.Background(), "czanel", "nazwa", 312, 342, 54)
	return s.persistence.InsertFilm(context.Background(), time.Now(), "czanel", "tytół", "descp", 3_000_000_000, 321, 123, 43)
}
