
-- name: CreateAccount :one
INSERT INTO account (
  id, owner, balance, currency
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: GetAllAccounts :many
SELECT * FROM account
LIMIT $1
OFFSET $2;


-- name: UpdateAccount :one
UPDATE account
  set balance = $2
WHERE id = $1
RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;