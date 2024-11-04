package youtube_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type YouTubeResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		SNIPPET struct {
			Title string `json:"title"`
		} `json:"snippet"`
	} `json:"items"`
}

// GetYoutubeVideoId_and_title делает запрос к YouTube API и возвращает ID видео
func GetYoutubeVideoId_and_title(apiKey string, query string) (string, string, error) {

	apiUrl := "https://www.googleapis.com/youtube/v3/search?part=snippet&type=video&q=" + url.QueryEscape(query) + "&key=" + apiKey

	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", "", fmt.Errorf("error occurred while fetching data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("YouTube API request failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read response body: %v", err)
	}

	var data YouTubeResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse JSON: %v", err)
	}

	if len(data.Items) > 0 && data.Items[0].ID.VideoID != "" {
		return data.Items[0].ID.VideoID, data.Items[0].SNIPPET.Title, nil
	}

	return "Video not found", "", nil
}
