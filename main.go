package main

import (
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/api/youtube/v3"
)

const DELIMITER = ","
const DIRECTORY_TO_STORE_OUTPUT = "/ytFiles"
const FILE_MODE = os.FileMode(0700)

const MAX_RESULTS = 50 // the max google allows rn is 50 per page.
const MINE = false

func main() {
	service := RunAuth()

	// decent auth check, can be refactored to yield channelID from user.
	// arguments := []string{"snippet", "contentDetails", "statistics"}
	// ChannelsListByUsername(service, arguments, "GoogleDevelopers")

	// fullParts := []string{"contentDetails","id","localizations","player","snippet","status"}

	err := godotenv.Load()
	if err != nil {
		HandleError("Error loading .env file", err)
	}

	channelId := os.Getenv("CHANNEL_ID")

	parts := []string{"contentDetails", "snippet"}

	nextPageToken := ""
	// need to call to find out how many items
	// I was going to use recursion or at least a while loop, but golang doesn't have null,
	// it just has zero case, which won't work for this specific implementation

	var playlists = []*youtube.Playlist{}
	var playlistListResponse *youtube.PlaylistListResponse

	for {
		// Retrieve next set of items in the playlist.
		playlistListResponse = PlaylistsListByChannelID(service, parts, channelId, MAX_RESULTS, nextPageToken)

		for _, playlist := range playlistListResponse.Items {
			playlists = append(playlists, playlist)
		}

		// Set the page token to retrieve the next page of results
		// exit the loop if all results have been retrieved.
		nextPageToken = playlistListResponse.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// fmt.Printf("does it work? %t", len(playlists) == int(playlistListResponse.PageInfo.TotalResults))
	// if (len(playlists) != int(playlistListResponse.PageInfo.TotalResults)) {
	// 	Warn("Please check, it seems we couldn't gather all playlist information", nil)
	// }

	var sb strings.Builder

	for _, playlist := range playlists {
		snippet := playlist.Snippet
		playlistTitle := snippet.Title

		sb.WriteString("PlayList: " + playlistTitle + "\n")
		for {
			// Retrieve next set of items in the playlist.
			itemsParts := []string{"snippet"}
			playlistResponse := PlaylistItemsList(service, itemsParts, playlist.Id, MAX_RESULTS, nextPageToken)

			for _, playlistItem := range playlistResponse.Items {
				itemSnippet := playlistItem.Snippet
				title := itemSnippet.Title
				videoId := itemSnippet.ResourceId.VideoId
				// description := itemSnippet.Description

				// fromChannelTitle := snippet.videoOwnerChannelTitle
				// fromChannelId := snippet.videoOwnerChannelId
				// sb.WriteString(title + DELIMITER + videoId + DELIMITER + fromChannelTitle + DELIMITER + fromChannelId + "\n")

				sb.WriteString(title + DELIMITER + videoId + "\n")
			}

			// Set the token to retrieve the next page of results
			// or exit the loop if all results have been retrieved.
			nextPageToken = playlistResponse.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
	}

	// Handle Output
	today := time.Now().Format("Jan02-2006") // golang magic number
	// could do t = time.Now()
	// fmt.Sprintf("%s %02d %d", time.Month().String()[:3], t.Day(), t.Year())

	fileName := "ytList" + today + ".md"

	byteArray := []byte(sb.String())
	homedir, err := os.UserHomeDir()
	fullDirPath := homedir + DIRECTORY_TO_STORE_OUTPUT
	os.MkdirAll(fullDirPath, FILE_MODE)
	err2 := os.WriteFile(fullDirPath+"/"+fileName, byteArray, FILE_MODE)
	HandleError("Did NOT write file", err2)
}
