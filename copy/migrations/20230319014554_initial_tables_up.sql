-- Add migration script here

CREATE TABLE events_copy(
  id uuid primary key,
  source inet not null,
  payload bytea,
  event_timestamp timestamp not null default now()
); 

CREATE INDEX on events_copy (source);

CREATE INDEX on events_copy (encode(payload, 'escape'));

CREATE INDEX on events_copy (event_timestamp);

CREATE TABLE events_insert(
  id uuid primary key,
  source inet not null,
  payload bytea,
  event_timestamp timestamp not null default now()
); 

CREATE INDEX on events_insert (source);

CREATE INDEX on events_insert (encode(payload, 'escape'));

CREATE INDEX on events_insert (event_timestamp);

