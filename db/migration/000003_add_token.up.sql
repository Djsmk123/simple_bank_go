CREATE TABLE tokens (
 "token" varchar PRIMARY KEY,
  "username" varchar NOT NULL,
  "issued_at" timestamptz NOT NULL,
  "expired_at" timestamptz NOT NULL
);