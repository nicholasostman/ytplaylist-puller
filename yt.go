package main

import (
	"fmt"

	"google.golang.org/api/youtube/v3"
)

func PlaylistsListByChannelID(service *youtube.Service, parts []string) {
	call := service.Playlists.List(parts)
	response, err := call.Do()
	HandleError("", err)
	fmt.Println(response)
}

func ChannelsListByUsername(service *youtube.Service, part []string, forUsername string) {
	call := service.Channels.List(part)
	call = call.ForUsername(forUsername)
	response, err := call.Do()
	HandleError("", err)
	fmt.Println(fmt.Sprintf("Channel's ID is %s, title is '%s', "+
		"and it has %d views.",
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].Statistics.ViewCount))
}
