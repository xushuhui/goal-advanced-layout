-- name: GetUser :one
SELECT * FROM user WHERE user_id = ? LIMIT 1;
-- name: GetUserByUsername :one
SELECT * FROM user WHERE username = ? LIMIT 1;
-- name: ListUser :many
SELECT * FROM user ORDER BY id DESC;
-- name: CreateUser :execresult
INSERT INTO user (user_id,username, password) VALUES (?,?, ?);

