CREATE TABLE yards (
  id SERIAL PRIMARY KEY,
  code TEXT UNIQUE NOT NULL,
  name TEXT
);

CREATE TABLE blocks (
  id SERIAL PRIMARY KEY,
  yard_id INT NOT NULL REFERENCES yards(id),
  code TEXT NOT NULL,
  rows INT NOT NULL,
  cols INT NOT NULL,
  max_tier INT NOT NULL DEFAULT 6,
  UNIQUE (yard_id, code)
);