-- TodoTable operations

-- name: GetTable :one
SELECT * FROM todo_table
WHERE id = ?;

-- name: GetTableByTitle :one
SELECT * FROM todo_table
WHERE title = ?;

-- name: ListTables :many
SELECT * FROM todo_table
ORDER BY title;

-- name: CreateTable :one
INSERT INTO todo_table (title)
VALUES (?)
RETURNING *;

-- name: UpdateTable :one
UPDATE todo_table
SET title = ?
WHERE id = ?
RETURNING *;

-- name: DeleteTable :exec
DELETE FROM todo_table
WHERE id = ?;

-- TodoEntry operations

-- name: GetEntry :one
SELECT * FROM todo_entry
WHERE id = ?;

-- name: ListEntries :many
SELECT * FROM todo_entry
WHERE table_id = ?
ORDER BY id;

-- name: ListAllEntries :many
SELECT * FROM todo_entry
ORDER BY table_id, id;

-- name: CreateEntry :one
INSERT INTO todo_entry (content, is_done, table_id)
VALUES (?, false, ?)
RETURNING *;

-- name: UpdateEntryContent :one
UPDATE todo_entry
SET content = ?
WHERE id = ?
RETURNING *;

-- name: UpdateEntryStatus :one
UPDATE todo_entry
SET is_done = ?
WHERE id = ?
RETURNING *;

-- name: ArchiveEntry :one
UPDATE todo_entry
SET is_archived = ?
WHERE id = ?
RETURNING *;

-- name: MoveEntry :one
UPDATE todo_entry
SET table_id = ?
WHERE id = ?
RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM todo_entry
WHERE id = ?;

-- Aggregate operations

-- name: CountTableEntries :one
SELECT COUNT(*) FROM todo_entry
WHERE table_id = ?;

-- name: ListEntriesByStatus :many
SELECT * FROM todo_entry
WHERE table_id = ? AND is_done = ?
ORDER BY id;
