package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

type YouTubeResponse struct {
	Items []struct {
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"snippet"`
	} `json:"items"`
}

func extractVideoID(link string) (string, error) {
	re := regexp.MustCompile(`(?:v=|youtu\.be/)([a-zA-Z0-9_-]{11})`)
	matches := re.FindStringSubmatch(link)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", fmt.Errorf("could not extract video ID")
}

func fetchVideoDetails(videoID, apiKey string) (*YouTubeResponse, error) {
	apiURL := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?part=snippet&id=%s&key=%s", videoID, apiKey)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ytResp YouTubeResponse
	if err := json.NewDecoder(resp.Body).Decode(&ytResp); err != nil {
		return nil, err
	}
	return &ytResp, nil
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		fmt.Println("API key not found in environment variables.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter YouTube video URL: ")
	link, _ := reader.ReadString('\n')
	link = strings.TrimSpace(link)

	videoID, err := extractVideoID(link)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	details, err := fetchVideoDetails(videoID, apiKey)
	if err != nil {
		fmt.Println("Failed to fetch video details:", err)
		return
	}

	if len(details.Items) == 0 {
		fmt.Println("No video found.")
		return
	}

	video := details.Items[0].Snippet
	fmt.Println("Title:", video.Title)
	fmt.Println("Description:", video.Description)
}
