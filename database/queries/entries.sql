-- name: GetEntriesFromLogbook :many
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
OFFSET $3;

-- name: GetEntryFromLogbook :one
SELECT
EntryId,
Title,
Description,
CreatedOn,
CreatedBy,
LogbookId
FROM entries 
WHERE EntryId=$1
AND LogbookId=$2;

-- name: GetLogbookAndOwnerFromEntry :one
SELECT 
entries.LogbookId,
logbooks.OwnedBy 
FROM logbooks 
JOIN entries 
ON entries.LogbookId = logbooks.LogbookId
WHERE EntryId=$1;

-- name: InsertNewEntryInLogbook :execresult
INSERT INTO entries(Title,Description,CreatedBy,LogbookId) VALUES
($1,$2,$3,$4);

-- name: DeleteEntryFromLogbook :execresult
DELETE FROM entries WHERE EntryId=$1 AND LogbookId=$2;

-- name: UpdateEntryFromLogbook :execresult
UPDATE entries 
SET title = $3,
description = $4
WHERE entryid = $1
AND logbookid = $2;
