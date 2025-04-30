package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	green = "\033[32m"
	reset = "\033[0m"
)

type YouTubeResponse struct {
	Items []struct {
		Snippet struct {
			Title        string    `json:"title"`
			Description  string    `json:"description"`
			PublishedAt  time.Time `json:"publishedAt"`
			ChannelTitle string    `json:"channelTitle"`
			Tags         []string  `json:"tags"`
		} `json:"snippet"`
		Statistics struct {
			ViewCount    string `json:"viewCount"`
			LikeCount    string `json:"likeCount"`
			CommentCount string `json:"commentCount"`
		} `json:"statistics"`
		ContentDetails struct {
			Duration        string `json:"duration"`
			Definition      string `json:"definition"`
			Caption         string `json:"caption"`
			LicensedContent bool   `json:"licensedContent"`
		} `json:"contentDetails"`
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
	apiURL := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?part=snippet,statistics,contentDetails&id=%s&key=%s", videoID, apiKey)
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

func printColorOutput(video YouTubeResponse) {
	item := video.Items[0]
	fmt.Println(green + "Title:       " + reset + item.Snippet.Title)
	fmt.Println(green + "Channel:     " + reset + item.Snippet.ChannelTitle)
	fmt.Println(green + "Published:   " + reset + item.Snippet.PublishedAt.Format("Jan 02, 2006 15:04"))
	fmt.Println(green + "Description: " + reset + item.Snippet.Description)
	fmt.Println(green + "Tags:        " + reset + strings.Join(item.Snippet.Tags, ", "))
	fmt.Println(green + "Duration:    " + reset + item.ContentDetails.Duration)
	fmt.Println(green + "Definition:  " + reset + strings.ToUpper(item.ContentDetails.Definition))
	fmt.Println(green + "Has Caption: " + reset + item.ContentDetails.Caption)
	fmt.Println(green + "Licensed:    " + reset + fmt.Sprintf("%t", item.ContentDetails.LicensedContent))
	fmt.Println(green + "Views:       " + reset + item.Statistics.ViewCount)
	fmt.Println(green + "Likes:       " + reset + item.Statistics.LikeCount)
	fmt.Println(green + "Comments:    " + reset + item.Statistics.CommentCount)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		fmt.Println("Missing YOUTUBE_API_KEY in .env")
		return
	}

	url := flag.String("url", "", "YouTube video URL")
	jsonOut := flag.Bool("json", false, "Print output in JSON format")
	flag.Parse()

	if *url == "" {
		fmt.Println("Usage: go run main.go -url <YouTube video URL> [--json]")
		return
	}

	videoID, err := extractVideoID(*url)
	if err != nil {
		fmt.Println("Error extracting video ID:", err)
		return
	}

	data, err := fetchVideoDetails(videoID, apiKey)
	if err != nil {
		fmt.Println("Failed to fetch data:", err)
		return
	}

	if len(data.Items) == 0 {
		fmt.Println("No video found.")
		return
	}

	if *jsonOut {
		jsonData, err := json.MarshalIndent(data.Items[0], "", "  ")
		if err != nil {
			fmt.Println("Error formatting JSON:", err)
			return
		}
		fmt.Println(string(jsonData))
	} else {
		printColorOutput(*data)
	}
}
