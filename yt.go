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

func PlaylistsListByChannelID(service *youtube.Service, parts []string, channelId string, maxResults int64, pageToken string) *youtube.PlaylistListResponse {
	call := service.Playlists.List(parts)
	call.ChannelId(channelId)
	call.MaxResults(maxResults)
	if pageToken != "" {
		call = call.PageToken(pageToken)
	}
	response, err := call.Do()
	HandleError("", err)
	return response
}

// Retrieve playlistItems (videos) in the specified playlist
func PlaylistItemsList(service *youtube.Service, parts []string, playlistId string, maxResults int64, pageToken string) *youtube.PlaylistItemListResponse {
	call := service.PlaylistItems.List(parts)
	call = call.PlaylistId(playlistId)
	call.MaxResults(maxResults)
	if pageToken != "" {
		call = call.PageToken(pageToken)
	}
	response, err := call.Do()
	HandleError("", err)
	return response
}

// Retrieve resource for the authenticated user's channel
func ChannelsListMine(service *youtube.Service, parts []string) *youtube.ChannelListResponse {
	call := service.Channels.List(parts)
	call = call.Mine(true)
	response, err := call.Do()
	HandleError("", err)
	return response
}
