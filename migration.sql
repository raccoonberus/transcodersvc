CREATE TABLE resource (
  id                SERIAL PRIMARY KEY,
  created_at        TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
  updated_at        TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
  removed_at        TIMESTAMPTZ            DEFAULT NULL,
  size              INTEGER       NOT NULL,
  original_filename VARCHAR(2048) NOT NULL,
  internal_filename VARCHAR(2048) NOT NULL
);

CREATE TABLE tasks (
  id                SERIAL PRIMARY KEY,
  created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  removed_at        TIMESTAMPTZ          DEFAULT NULL,
  resource_id       INTEGER     NOT NULL REFERENCES resource (id),
  callback_url      VARCHAR(2048)        DEFAULT NULL,
  callback_response TEXT                 DEFAULT NULL,
  started_at        TIMESTAMPTZ          DEFAULT NULL,
  finished_at       TIMESTAMPTZ          DEFAULT NULL,
  output            TEXT                 DEFAULT NULL
);