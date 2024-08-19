-- name: GetLogbooksOwnedBy :many
SELECT owned_by.LogbookId FROM owned_by NATURAL JOIN users WHERE owned_by.UserId=$1;
