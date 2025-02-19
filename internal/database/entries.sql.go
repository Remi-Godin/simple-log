// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: entries.sql

package database

import (
	"context"
	"database/sql"
)

const deleteEntryFromLogbook = `-- name: DeleteEntryFromLogbook :execresult
DELETE FROM entries WHERE EntryId=$1 AND LogbookId=$2
`

type DeleteEntryFromLogbookParams struct {
	Entryid   int32
	Logbookid int32
}

func (q *Queries) DeleteEntryFromLogbook(ctx context.Context, arg DeleteEntryFromLogbookParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteEntryFromLogbook, arg.Entryid, arg.Logbookid)
}

const getEntriesFromLogbook = `-- name: GetEntriesFromLogbook :many
SELECT
EntryId,
Title,
Description,
CreatedOn,
CreatedBy,
LogbookId
FROM entries 
WHERE LogbookId=$1
ORDER BY CreatedOn DESC
LIMIT $2
OFFSET $3
`

type GetEntriesFromLogbookParams struct {
	Logbookid int32
	Limit     int32
	Offset    int32
}

func (q *Queries) GetEntriesFromLogbook(ctx context.Context, arg GetEntriesFromLogbookParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, getEntriesFromLogbook, arg.Logbookid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.Entryid,
			&i.Title,
			&i.Description,
			&i.Createdon,
			&i.Createdby,
			&i.Logbookid,
		); err != nil {
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

const getEntryFromLogbook = `-- name: GetEntryFromLogbook :one
SELECT
EntryId,
Title,
Description,
CreatedOn,
CreatedBy,
LogbookId
FROM entries 
WHERE EntryId=$1
AND LogbookId=$2
`

type GetEntryFromLogbookParams struct {
	Entryid   int32
	Logbookid int32
}

func (q *Queries) GetEntryFromLogbook(ctx context.Context, arg GetEntryFromLogbookParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntryFromLogbook, arg.Entryid, arg.Logbookid)
	var i Entry
	err := row.Scan(
		&i.Entryid,
		&i.Title,
		&i.Description,
		&i.Createdon,
		&i.Createdby,
		&i.Logbookid,
	)
	return i, err
}

const getLogbookAndOwnerFromEntry = `-- name: GetLogbookAndOwnerFromEntry :one
SELECT 
entries.LogbookId,
logbooks.OwnedBy 
FROM logbooks 
JOIN entries 
ON entries.LogbookId = logbooks.LogbookId
WHERE EntryId=$1
`

type GetLogbookAndOwnerFromEntryRow struct {
	Logbookid int32
	Ownedby   string
}

func (q *Queries) GetLogbookAndOwnerFromEntry(ctx context.Context, entryid int32) (GetLogbookAndOwnerFromEntryRow, error) {
	row := q.db.QueryRowContext(ctx, getLogbookAndOwnerFromEntry, entryid)
	var i GetLogbookAndOwnerFromEntryRow
	err := row.Scan(&i.Logbookid, &i.Ownedby)
	return i, err
}

const insertNewEntryInLogbook = `-- name: InsertNewEntryInLogbook :execresult
INSERT INTO entries(Title,Description,CreatedBy,LogbookId) VALUES
($1,$2,$3,$4)
`

type InsertNewEntryInLogbookParams struct {
	Title       string
	Description string
	Createdby   string
	Logbookid   int32
}

func (q *Queries) InsertNewEntryInLogbook(ctx context.Context, arg InsertNewEntryInLogbookParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertNewEntryInLogbook,
		arg.Title,
		arg.Description,
		arg.Createdby,
		arg.Logbookid,
	)
}

const updateEntryFromLogbook = `-- name: UpdateEntryFromLogbook :execresult
UPDATE entries 
SET title = $3,
description = $4
WHERE entryid = $1
AND logbookid = $2
`

type UpdateEntryFromLogbookParams struct {
	Entryid     int32
	Logbookid   int32
	Title       string
	Description string
}

func (q *Queries) UpdateEntryFromLogbook(ctx context.Context, arg UpdateEntryFromLogbookParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateEntryFromLogbook,
		arg.Entryid,
		arg.Logbookid,
		arg.Title,
		arg.Description,
	)
}
