-- name: InsertNewUser :execresult
INSERT INTO users(FirstName,LastName,Email,PasswordHash) VALUES
($1,$2,$3,$4);

-- name: GetUserInfo :one
SELECT
FirstName,
LastName
FROM
users
WHERE
Email = $1;

-- name: GetUserPasswordHash :one
SELECT
PasswordHash
FROM
users
WHERE
Email = $1;
