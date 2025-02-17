-- name: CreateBank :one
INSERT INTO banks (
    id,
    planet_id,
    organization_id,
    name
)
VALUES (
    @id,
    @planet_id,
    @organization_id,
    @name
)
RETURNING id;


-- name: UpdateBankName :exec
UPDATE banks SET
    name = @name
WHERE id = @id;


-- name: DeleteBank :exec
DELETE FROM
    banks
WHERE
    id = @id;


-- name: SelectBankByID :one
SELECT
    id,
    planet_id,
    organization_id,
    name
FROM
    banks
WHERE
    id = @id;


-- name: SelectBanksByPlanetID :many
SELECT
    id,
    planet_id,
    organization_id,
    name
FROM
    banks
WHERE
    planet_id = @planet_id;