package persistence

import "time"

type Channel struct {
	Id    string `db:"channel_id"`
	Title string `db:"title"`
}

type Film struct {
	Id           string    `db:"id"`
	PublishedAt  time.Time `db:"published_at"`
	ChannelID    string    `db:"channel_id"`
	Title        string    `db:"title"`
	Description  string    `db:"descript"`
	ViewCount    int64     `db:"view_count"`
	LikeCount    int32     `db:"like_count"`
	DislikeCount int32     `db:"dislike_count"`
	CommentCount int32     `db:"comment_count"`
}
