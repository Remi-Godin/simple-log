-- name: GetLogbooksOwnedBy :many
SELECT 
l.LogbookId,
l.Title,
l.Description,
l.CreatedOn ,
u.FirstName,
u.LastName,
u.Email
FROM logbooks l
INNER JOIN users u 
ON l.OwnedBy = u.Email
WHERE l.OwnedBy=$1
ORDER BY CreatedOn DESC
LIMIT $2
OFFSET $3;

-- name: GetLogbooks :many
SELECT
LogbookId,
Title,
OwnedBy
FROM logbooks
LIMIT $1
OFFSET $2;

-- name: GetLogbookData :one
SELECT
LogbookId,
Title,
Description,
OwnedBy
FROM logbooks
WHERE LogbookId = $1;

-- name: InsertNewLogbook :execresult
INSERT INTO logbooks(Title,Description,OwnedBy) VALUES
($1,$2,$3);

-- name: DeleteLogbook :execresult
DELETE FROM logbooks WHERE LogbookId=$1;
