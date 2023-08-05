-- name: CreateToken :one
INSERT INTO tokens (
  token,
  username,
  issued_at,
  expired_at
) VALUES (
  $1, $2, $3,$4
) RETURNING *;

-- name: GetToken :one
SELECT * FROM tokens
WHERE username= $1 AND expired_at > Now() LIMIT 1;


-- name: UpdateToken :one
UPDATE tokens
SET expired_at = Now()
WHERE username = $1
RETURNING *;

