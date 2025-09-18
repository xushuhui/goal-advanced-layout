-- name: GetUser :one
SELECT * FROM user WHERE user_id = ? AND deleted_at IS NULL LIMIT 1;
-- name: GetUserByUsername :one
SELECT * FROM user WHERE username = ? AND deleted_at IS NULL LIMIT 1;
-- name: ListUser :many
SELECT * FROM user WHERE deleted_at IS NULL ORDER BY id DESC;
-- name: CreateUser :execresult
INSERT INTO user (user_id,username, password) VALUES (?,?,?);
-- name: DeleteUser :exec
UPDATE user SET deleted_at=NOW() WHERE user_id = ?;
-- name: UpdateUser :exec
UPDATE user SET username = ?, password = ? WHERE user_id = ?;
-- name: GetUsersByPage :many
SELECT * FROM user WHERE deleted_at IS NULL ORDER BY id DESC LIMIT ? OFFSET ?;
