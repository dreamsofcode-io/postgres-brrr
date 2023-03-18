PREPARE test_insert (text) AS
  INSERT INTO users (name) VALUES ($1);

EXECUTE test_insert('User ' || :pgbench_tid);

