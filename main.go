package main

import (
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const DELIMITER = ","
const DIRECTORY_TO_STORE_OUTPUT = "/ytFiles"
const FILE_MODE = os.FileMode(0700)

func main() {
	service := RunAuth()

	// decent auth check, can be refactored to yield channelID from user.
	// arguments := []string{"snippet", "contentDetails", "statistics"}
	// ChannelsListByUsername(service, arguments, "GoogleDevelopers")

	// fullParts := []string{"contentDetails","id","localizations","player","snippet","status"}

	// parts := []string{"snippet", "id=%s", os.Getenv("CHANNEL_ID")}

	err := godotenv.Load()
	if err != nil {
		HandleError("Error loading .env file", err)
	}

	channelId := os.Getenv("CHANNEL_ID")

	parts := []string{"contentDetails", "snippet"}
	playlistList := PlaylistsListByChannelID(service, parts, channelId)

	today := time.Now().Format("Jan02-2006") // golang magic number
	// could do t = time.Now()
	// fmt.Sprintf("%s %02d %d", time.Month().String()[:3], t.Day(), t.Year())

	fileName := "ytList" + today + ".md"

	var sb strings.Builder

	// this could break if massive number of playlists etc.
	// could be better, but MVP first
	// for _, playlist := range playlistList.Items {
	// 	playlistId := playlist.Id
	// 	snippet := playlist.Snippet
	// 	playlistTitle := snippet.Title
	// 	// desc := snippet.Description

	// 	// Print the playlist ID and title for the playlist resource.
	// 	fmt.Println(playlistId, ": ", playlistTitle)
	// 	sb.WriteString(playlistTitle)
	// }

	// maxResults := 50 // the max google allows rn is 50 per page.

	nextPageToken := ""
	// I may refactor this, will think after I sleep lol, but could put in to have all pages be default returned
	for _, playlist := range playlistList.Items {

		sb.WriteString("PlayList: " + playlist.Snippet.Title + "\n")
		for {
			// Retrieve next set of items in the playlist.
			itemsParts := []string{"snippet"}
			playlistResponse := playlistItemsList(service, itemsParts, playlist.Id, nextPageToken)

			for _, playlistItem := range playlistResponse.Items {
				title := playlistItem.Snippet.Title
				videoId := playlistItem.Snippet.ResourceId.VideoId

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

	byteArray := []byte(sb.String())
	homedir, err := os.UserHomeDir()
	fullDirPath := homedir + DIRECTORY_TO_STORE_OUTPUT
	os.MkdirAll(fullDirPath, FILE_MODE)
	err2 := os.WriteFile(fullDirPath+"/"+fileName, byteArray, FILE_MODE)
	HandleError("Did NOT write file", err2)
}
