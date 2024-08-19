-- name: GetAllEntriesFromLogbook :many
SELECT 
entries.EntryId,
entries.Title,
entries.Description,
entries.CreatedOn,
entries.CreatedBy 
FROM entries 
NATURAL JOIN belongs_to 
WHERE belongs_to.LogbookId=$1;

-- name: GetLastNEntriesFromLogbook :many
SELECT 
entries.EntryId,
entries.Title,
entries.Description,
entries.CreatedOn,
entries.CreatedBy 
FROM entries 
NATURAL JOIN belongs_to 
WHERE belongs_to.LogbookId=$1 
ORDER BY entries.CreatedOn 
LIMIT $2 
OFFSET $3;

-- name: GetEntryFromLogbook :one
SELECT 
entries.EntryId,
entries.Title,
entries.Description,
entries.CreatedOn,
entries.CreatedBy 
FROM entries 
NATURAL JOIN belongs_to
WHERE belongs_to.LogbookId=$1 AND entries.EntryId=$2;
