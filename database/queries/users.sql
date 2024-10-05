-- name: InsertNewUser :execresult
INSERT INTO users(FirstName,LastName,Email,PasswordHash) VALUES
($1,$2,$3,$4);

-- name: GetUserInfo :one
SELECT
FirstName,
LastName,
Email
FROM
users
WHERE
userid = $1;

-- name: GetUserPasswordHash :one
SELECT
Email,
PasswordHash
FROM
users
WHERE
userid = $1;

-- name: GetUserInfoFromEmail :one
SELECT
FIrstName,
LastName
FROM
users
WHERE
Email = $1;
