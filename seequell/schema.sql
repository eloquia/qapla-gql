CREATE TABLE users (
  id          BIGSERIAL,
  first_name  VARCHAR(64)   NOT NULL,
  last_name   VARCHAR(64)   NOT NULL,
  middle_name VARCHAR(64),
  goes_by     VARCHAR(64),
  email       VARCHAR(128)  NOT NULL UNIQUE,
  pw          VARCHAR(128)  NOT NULL,
  gender      VARCHAR(64),
  ethnicity   VARCHAR(64),
  position    VARCHAR(64),
  institution VARCHAR(64),
  created_at  timestamp     NOT NULL DEFAULT NOW(),
  updated_at  timestamp,
  PRIMARY KEY(id)
);

CREATE TABLE projects (
  id            BIGSERIAL     PRIMARY KEY,
  project_name  VARCHAR(64)   NOT NULL,
  project_desc  TEXT,
  created_at    timestamp     NOT NULL DEFAULT NOW(),
  updated_at    timestamp
);

CREATE TABLE project_user (
  user_id BIGSERIAL,
  project_id BIGSERIAL,
  PRIMARY KEY (user_id, project_id),
  CONSTRAINT project_user_fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT project_user_fk_project FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);
