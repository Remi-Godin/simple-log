// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: logbooks.sql

package database

import (
	"context"
)

const getLogbooksOwnedBy = `-- name: GetLogbooksOwnedBy :many
SELECT owned_by.LogbookId FROM owned_by NATURAL JOIN users WHERE owned_by.UserId=$1
`

func (q *Queries) GetLogbooksOwnedBy(ctx context.Context, userid int32) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, getLogbooksOwnedBy, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var logbookid int32
		if err := rows.Scan(&logbookid); err != nil {
			return nil, err
		}
		items = append(items, logbookid)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
