-- name: GetLogbooksOwnedBy :many
SELECT 
LogbookId,
Title,
OwnedBy 
FROM logbooks 
WHERE OwnedBy=$1;

-- name: GetLogbooks :many
SELECT
LogbookId,
Title,
OwnedBy
FROM logbooks
LIMIT $1
OFFSET $2;

-- name: InsertNewLogbook :execresult
INSERT INTO logbooks(Title,OwnedBy) VALUES
($1,$2);

-- name: DeleteLogbook :execresult
DELETE FROM logbooks WHERE LogbookId=$1;
