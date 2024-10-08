// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: logbooks.sql

package database

import (
	"context"
	"database/sql"
)

const deleteLogbook = `-- name: DeleteLogbook :execresult
DELETE FROM logbooks WHERE LogbookId=$1
`

func (q *Queries) DeleteLogbook(ctx context.Context, logbookid int32) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteLogbook, logbookid)
}

const getLogbooks = `-- name: GetLogbooks :many
SELECT
LogbookId,
Title,
OwnedBy
FROM logbooks
LIMIT $1
OFFSET $2
`

type GetLogbooksParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetLogbooks(ctx context.Context, arg GetLogbooksParams) ([]Logbook, error) {
	rows, err := q.db.QueryContext(ctx, getLogbooks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Logbook
	for rows.Next() {
		var i Logbook
		if err := rows.Scan(&i.Logbookid, &i.Title, &i.Ownedby); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogbooksOwnedBy = `-- name: GetLogbooksOwnedBy :many
SELECT 
LogbookId,
Title,
OwnedBy 
FROM logbooks 
WHERE OwnedBy=$1
`

func (q *Queries) GetLogbooksOwnedBy(ctx context.Context, ownedby int32) ([]Logbook, error) {
	rows, err := q.db.QueryContext(ctx, getLogbooksOwnedBy, ownedby)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Logbook
	for rows.Next() {
		var i Logbook
		if err := rows.Scan(&i.Logbookid, &i.Title, &i.Ownedby); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertNewLogbook = `-- name: InsertNewLogbook :execresult
INSERT INTO logbooks(Title,OwnedBy) VALUES
($1,$2)
`

type InsertNewLogbookParams struct {
	Title   string
	Ownedby int32
}

func (q *Queries) InsertNewLogbook(ctx context.Context, arg InsertNewLogbookParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertNewLogbook, arg.Title, arg.Ownedby)
}
