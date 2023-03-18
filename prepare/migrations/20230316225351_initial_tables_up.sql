-- Add migration script here
create table task (
  id SERIAL PRIMARY KEY,
  name varchar not null,
  due_date date not null,
  is_done boolean not null default false,
  created_at timestamp NOT NULL DEFAULT now()
)
