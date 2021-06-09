package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	firstPageToken = ""
	lastPageToken  = ""
)

func GetAllFilmsAfter(apiURL, apiKey, channelID string, publishedAfter time.Time) ([]ListResponse, error) {
	allFilms := []ListResponse{}
	singleListResponse, err := GetListResponse(apiURL, apiKey, channelID, firstPageToken, publishedAfter)
	if err != nil {
		return allFilms, err
	}

	allFilms = append(allFilms, singleListResponse)

	for singleListResponse.NextPageToken != lastPageToken {
		singleListResponse, err = GetListResponse(apiURL, apiKey, channelID, singleListResponse.NextPageToken, publishedAfter)
		if err != nil {
			return allFilms, err
		}
		allFilms = append(allFilms, singleListResponse)
	}
	return allFilms, nil
}

func GetListResponse(apiURL, apiKey, channelID, pageToken string, publishedAfter time.Time) (ListResponse, error) {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return ListResponse{}, err
	}

	q := req.URL.Query()
	q.Add("part", "snippet")
	q.Add("order", "date")
	q.Add("type", "video")
	q.Add("maxResults", "50")
	if pageToken != "" {
		q.Add("pageToken", pageToken)
	}
	q.Add("publishedAfter", publishedAfter.Format(time.RFC3339))
	q.Add("key", apiKey)
	q.Add("channelId", channelID)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return ListResponse{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ListResponse{}, err
	}

	listResponse := ListResponse{}
	err = json.Unmarshal(body, &listResponse)
	if err != nil {
		return ListResponse{}, err
	}

	return listResponse, nil
}
