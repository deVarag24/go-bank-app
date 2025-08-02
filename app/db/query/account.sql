
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

-- name: GetAccountForUpdate :one
SELECT * FROM account
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: GetAllAccounts :many
SELECT * FROM account
LIMIT $1
OFFSET $2;


-- name: UpdateAccountBalance :one
UPDATE account
  set balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE account
  set balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;