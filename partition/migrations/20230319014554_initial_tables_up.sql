-- Add migration script here
CREATE TABLE events (
  id uuid,
  source inet not null,
  payload bytea,
  event_timestamp timestamp not null default now(),
  primary key (id, event_timestamp)
) PARTITION BY RANGE (event_timestamp);

CREATE TABLE events_2023_03_16 PARTITION OF events
  FOR VALUES FROM ('2023-03-16') TO ('2023-03-17');

CREATE TABLE events_2023_03_17 PARTITION OF events
  FOR VALUES FROM ('2023-03-17') TO ('2023-03-18');

CREATE TABLE events_2023_03_18 PARTITION OF events
  FOR VALUES FROM ('2023-03-18') TO ('2023-03-19');

CREATE INDEX on events (source);

CREATE INDEX on events (encode(payload, 'escape'));

CREATE TABLE events_full(
  id uuid primary key,
  source inet not null,
  payload bytea,
  event_timestamp timestamp not null default now()
); 

CREATE INDEX on events_full (source);

CREATE INDEX on events_full (encode(payload, 'escape'));

CREATE INDEX on events_full (event_timestamp);
