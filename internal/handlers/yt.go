package handlers

type ytClient struct {
	apiKey string
}

func NewYT(apiKey string) ytClient {
	return ytClient{
		apiKey: apiKey,
	}
}
