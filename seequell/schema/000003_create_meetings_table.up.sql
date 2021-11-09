CREATE TABLE core_qapla.meetings (
  meeting_id BIGSERIAL,
  meeting_name  TEXT NOT NULL,
  created_at    timestamptz NOT NULL DEFAULT NOW(),
  updated_at    timestamptz,
  PRIMARY KEY (meeting_id)
);
