# YouTube Recon

**YouTube Recon** is a simple CLI tool written in Go that fetches metadata from any YouTube video — including title, description, channel info, tags, view count, and more. You can also optionally export the data in JSON format.

## Features

- Fetches:
  - Title
  - Channel Name
  - Description
  - Tags
  - Published Date
  - Views, Likes, Comments
  - Duration, Definition, Licensing
- Pretty CLI output with color
- JSON output with `--json` flag
- Secure use of API key via `.env`

---

## Prerequisites

- Go installed: [https://go.dev/dl/](https://go.dev/dl/)
- YouTube Data API v3 key from: [https://console.cloud.google.com/apis](https://console.cloud.google.com/apis)

---

## Setup

1. Clone the repository:

```bash
git clone https://github.com/yourusername/youtube-recon.git
cd youtube-recon```

2. Create a .env file:

```bash
YOUTUBE_API_KEY=your_api_key_here```

3. Run the tool:
```bash
go run main.go -url "https://www.youtube.com/watch?v=VIDEO_ID"```

For JSON output:
```bash
go run main.go -url "https://www.youtube.com/watch?v=VIDEO_ID" --json```

---

### Example Output

#### Color Output

Title:       Understanding Cybersecurity
Channel:     Tech Talks
Published:   Apr 12, 2024 18:30
Description: A quick walkthrough on the basics of cybersecurity...
Tags:        cyber, security, basics, intro
Duration:    PT8M14S
Definition:  HD
Has Caption: true
Licensed:    true
Views:       12645
Likes:       532
Comments:    23

#### JSON Output
```json
{
  "snippet": {
    "title": "Understanding Cybersecurity",
    "description": "A quick walkthrough on the basics...",
    "channelTitle": "Tech Talks",
    ...
  },
  "statistics": {
    "viewCount": "12645",
    "likeCount": "532",
    ...
  }
}```

---

## License

MIT License


---

## Author

Developed by [abdulConsole](https://abdulconsole.github.io)
