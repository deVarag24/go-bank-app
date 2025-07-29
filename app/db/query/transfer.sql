-- name: CreateTransfer :one
INSERT INTO transfer (
  id, from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfer
WHERE id = $1 LIMIT 1;

-- name: GetAllTransferByAccounts :many
SELECT * FROM transfer
WHERE from_account_id = $1
OR to_account_id = $2
LIMIT $3
OFFSET $4;