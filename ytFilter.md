// package main

// // I was hoping to setup a reusable filters struct, but
// // each service is specifically typed and calls are
// // var call *youtube.PlaylistsListCall which is what my struct data would get thrown on...
// // basically I'd have to implement this all differently so commented out for now
// // maybe as a *query?

// type YtFilter struct {
// 	channelId string
// 	// hl string
// 	maxResults int64
// 	mine bool
// 	// onBehalfOfContentOwner string // I just don't need this
// 	pageToken string
// 	playlistId string
// }

// func NewYtFilter( maxResults int64, mine bool, pageToken string) *YtFilter {
// 	YtFilter{maxResults: int64, mine: bool, pageToken: string}

// 	// default is en_US
// 	// if hl != "" {
// 	// 	YtFilter.hl= os.Getenv("LANG_CODE")
// 	// }

// 	if maxResults != 0 {
// 		YtFilter.maxResults = maxResults
// 	} else {
// 		YtFilter.maxResults = 50
// 	}

// 	if mine != false {
// 		YtFilter.mine = true
// 	}

// 	if pageToken != "" {
// 		YtFilter.pageToken = pageToken
// 	}
// 	return &YtFilter
// }

// // function InitServiceCall(maxResults int64, mine bool, pageToken string){
// // call = call.MaxResults(maxResults)
// // if mine != false {
// // 	call = call.Mine(true)
// // }
// // if pageToken != "" {
// // 	call = call.PageToken(pageToken)
// // }
// // }