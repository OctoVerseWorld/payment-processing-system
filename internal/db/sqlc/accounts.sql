-- name: SelectAccount :one
SELECT
    id,
    owner_id,
    owner_type,
    bank_id,
    currency_id,
    balance,
    is_reserve
FROM
    accounts
WHERE
    id = @id
    LIMIT 1;
