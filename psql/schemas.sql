-- Example queries for sqlc
CREATE TABLE tasks (
  id   BIGSERIAL PRIMARY KEY,
  done boolean      NOT NULL,
  description  text NOT NULL,
  user_id bigint not null,
  created_at timestamp not null
);


CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  username text not null,
  fullname text not null
);

CREATE TABLE events (
  id   BIGSERIAL PRIMARY KEY,
  user_id bigint not null,
  description text not null,
  scheduled_for timestamp not null,
  created_at timestamp not null
);

CREATE TABLE notes (
  id   BIGSERIAL PRIMARY KEY,
  user_id bigint not null,
  description text not null,
  created_at timestamp not null
);
