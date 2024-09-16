package main

import (
	"fmt"

	"google.golang.org/api/youtube/v3"
)

func ChannelsListByUsername(service *youtube.Service, parts []string, forUsername string) {
	call := service.Channels.List(parts)
	call = call.ForUsername(forUsername)
	response, err := call.Do()
	HandleError("", err)
	fmt.Println(fmt.Sprintf("Channel's ID is %s, title is '%s', "+
		"and it has %d views.",
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].Statistics.ViewCount))
}

// func PlaylistsListByChannelID(service *youtube.Service, parts []string, channelId string, pageToken string) *youtube.PlaylistListResponse {
func PlaylistsListByChannelID(service *youtube.Service, parts []string, channelId string) *youtube.PlaylistListResponse {
	call := service.Playlists.List(parts)
	call.ChannelId(channelId)
	// if pageToken != "" {
	// 	call = call.PageToken(pageToken)
	// }
	response, err := call.Do()
	HandleError("", err)
	return response
}

// Retrieve playlistItems in the specified playlist
func playlistItemsList(service *youtube.Service, parts []string, playlistId string, pageToken string) *youtube.PlaylistItemListResponse {
	call := service.PlaylistItems.List(parts)
	call = call.PlaylistId(playlistId)
	if pageToken != "" {
		call = call.PageToken(pageToken)
	}
	response, err := call.Do()
	HandleError("", err)
	return response
}
