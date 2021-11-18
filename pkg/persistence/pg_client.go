package persistence

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type pgClient struct {
	*sqlx.DB
}

func NewPostgresClient(db *sqlx.DB) pgClient {
	return pgClient{
		DB: db,
	}
}

func (p pgClient) InsertChannel(
	ctx context.Context,
	id, title string,
	viewCount int64,
	subscriberCount, videoCount int32,
) error {
	q := `
		INSERT INTO channels (channel_id, title, view_count, subscriber_count, video_count) 
			VALUES ($1, $2, $3, $4, $5)
	`
	_, err := p.ExecContext(ctx, q, id, title, viewCount, subscriberCount, videoCount)
	if err != nil {
		log.Error().Err(err).Msg("cannot insert channel to db")
		return err
	}
	return nil
}

func (p pgClient) GetChannel(ctx context.Context, id string) (Channel, error) {
	q := "SELECT channel_id, title FROM channels WHERE channel_id = $1"
	ch := Channel{}
	err := p.Get(&ch, q, id)
	if err != nil {
		if err != sql.ErrNoRows {
			return Channel{}, err
		}
		log.Info().Str("channel_id", id).Msg("channel not found")
		return Channel{}, nil
	}
	return ch, nil
}

func (p pgClient) InsertFilm(
	ctx context.Context,
	publishedAt time.Time,
	channelId, title, description string,
	viewCount int64,
	likeCount, dislikeCount, commentCount int32,
) error {
	q := `
		INSERT INTO films (published_at, channel_id, title, descript, view_count, like_count, dislike_count, comment_count) 
		VALUES ($1, $2, $3 , $4, $5, $6, $7, $8)
	`
	_, err := p.ExecContext(
		ctx,
		q,
		publishedAt,
		channelId,
		title,
		description,
		viewCount,
		likeCount,
		dislikeCount,
		commentCount,
	)
	if err != nil {
		log.Error().Err(err).Msg("cannot insert film to db")
		return err
	}
	return nil
}

func (p pgClient) GetFilmsForChannel(ctx context.Context, channelId string) ([]Film, error) {
	q := "SELECT id, published_at, channel_id, title, descript FROM films WHERE channel_id = $1"
	films := []Film{}
	err := p.SelectContext(ctx, &films, q, channelId)
	if err != nil {
		return nil, err
	}
	return films, nil
}
