// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: commands.sql

package queries

import (
	"context"
)

const getAnswer = `-- name: GetAnswer :one
SELECT answer FROM commands
         WHERE channel_id = $1
         AND command = $2
`

type GetAnswerParams struct {
	ChannelID int32  `json:"channel_id"`
	Command   string `json:"command"`
}

func (q *Queries) GetAnswer(ctx context.Context, arg GetAnswerParams) (string, error) {
	row := q.db.QueryRow(ctx, getAnswer, arg.ChannelID, arg.Command)
	var answer string
	err := row.Scan(&answer)
	return answer, err
}