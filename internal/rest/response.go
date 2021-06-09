package rest

import "time"

type ListResponse struct {
	Kind          string   `json:"kind"`
	Etag          string   `json:"etag"`
	PrevPageToken string   `json:"prevPageToken"`
	NextPageToken string   `json:"nextPageToken`
	RegionCode    string   `json:"regionCode"`
	PageInfo      PageInfo `json:"pageInfo"`
	Items         []Item   `json:"items"`
}

type PageInfo struct {
	TotalResults   int
	ResultsPerPage int
}

type Item struct {
	Kind    string
	Etag    string
	Id      Id
	Snippet Snippet
}

type Id struct {
	Kind    string
	VideoId string
}

type Snippet struct {
	PublishedAt          time.Time
	ChannelId            string
	Title                string
	Description          string
	Thumbnails           Thumbnails
	ChannelTitle         string
	LiveBroadcastContent string
	PublishTime          time.Time
}

type Thumbnails struct {
	Def    Thumbnail `json:"default"`
	Medium Thumbnail
	High   Thumbnail
}

type Thumbnail struct {
	Url    string
	Width  int
	Height int
}
