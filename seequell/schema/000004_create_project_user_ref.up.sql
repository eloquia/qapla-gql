-- Projects can have many users; users can be a part of many projects
CREATE TABLE core_qapla.project_users (
  project_id BIGSERIAL,
  user_id BIGSERIAL,
  PRIMARY KEY (project_id, user_id),
  CONSTRAINT project_users_project_fk FOREIGN KEY (project_id) REFERENCES core_qapla.projects(project_id),
  CONSTRAINT project_users_user_fk FOREIGN KEY (user_id) REFERENCES core_qapla.users(user_id)
);
