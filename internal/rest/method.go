package rest

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetListResponse(apiURL, apiKey, channelID string) (ListResponse, error) {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return ListResponse{}, err
	}

	q := req.URL.Query()
	q.Add("part", "snippet")
	q.Add("order", "date")
	q.Add("type", "video")
	q.Add("key", apiKey)
	q.Add("channelId", channelID)
	req.URL.RawQuery = q.Encode()

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
