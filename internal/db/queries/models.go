// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package queries

import ()

type Channel struct {
	ID         int32  `json:"id"`
	TwitchName string `json:"twitch_name"`
	TwitchID   int32  `json:"twitch_id"`
}

type Command struct {
	ID        int32 `json:"id"`
	ChannelID int32 `json:"channel_id"`
	// Move to another table, replace with tag_id
	Command string `json:"command"`
	// Move to another table, replace with answer_id
	Answer string `json:"answer"`
}
