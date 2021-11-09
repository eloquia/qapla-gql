CREATE TABLE core_qapla.projects (
  project_id    BIGSERIAL   NOT NULL,
  project_name  VARCHAR(64) NOT NULL,
  project_desc  TEXT,
  created_at    timestamptz NOT NULL DEFAULT NOW(),
  updated_at    timestamptz,
  PRIMARY KEY (project_id)
);

CREATE TABLE core_qapla.project_user (
  user_id BIGSERIAL,
  project_id BIGSERIAL,
  PRIMARY KEY (user_id, project_id),
  CONSTRAINT project_user_fk_user FOREIGN KEY (user_id) REFERENCES core_qapla.users(user_id),
  CONSTRAINT project_user_fk_project FOREIGN KEY (project_id) REFERENCES core_qapla.projects(project_id)
);
