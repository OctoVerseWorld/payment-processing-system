-- name: CreateAccount :one
INSERT INTO accounts (
    id,
    owner_id,
    currency_id,
    balance

)
VALUES (
    @id,
    @owner_id,
    @currency_id,
    @balance
)
RETURNING id;

-- name: UpdateAccount :one
UPDATE accounts SET
    owner_id = @owner_id,
    currency_id = @currency_id,
    balance = @balance
WHERE id = @id
RETURNING id AS res;

-- name: DeleteAccount :one
DELETE FROM
    accounts
WHERE
    id = @id
RETURNING id AS res;

-- name: SelectAccount :one
SELECT
    id,
    owner_id,
    currency_id,
    balance
FROM
    accounts
WHERE
    id = @id
LIMIT 1;