package persistance

import "time"

type Film struct {
	ID           string    `json:"id"`
	PublishedAt  time.Time `json:"published_at"`
	ChannelID    string    `json:"channel_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ChannelTitle string    `json:"channel_title"`
	PublishTime  time.Time `json:"publish_time"`
}
