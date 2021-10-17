-- name: ListTasks :many
SELECT * FROM tasks
WHERE user_id = $1 and created_at::date = $2::date;


-- name: CreateTask :one
INSERT INTO tasks (
  done, description, user_id, created_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks set done = $1 where user_id = $2 and id = $3
RETURNING *;


-- name: ListNotes :many
SELECT * FROM notes
WHERE user_id = $1 and created_at::date = $2::date;


-- name: CreateNote :one
INSERT INTO notes (
  description, user_id, created_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListEvents :many
SELECT * FROM events
WHERE user_id = $1 and created_at::date = $2::date;


-- name: CreateEvent :one
INSERT INTO events (
  description, user_id, created_at, scheduled_for
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;


-- name: CreateUser :one
INSERT INTO users (
  id, username, fullname
) VALUES (
  $1, $2, $3
)
RETURNING *;