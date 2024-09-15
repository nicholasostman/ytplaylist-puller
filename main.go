package main

func main() {
	// Hello() // to run hello world sanity check
	service := RunAuth()

	// decent auth check
	arguments := []string{"snippet", "contentDetails", "statistics"}
	ChannelsListByUsername(service, arguments, "GoogleDevelopers")

	// parts := []string{"id="}
	// PlaylistsListByChannelID()
}
