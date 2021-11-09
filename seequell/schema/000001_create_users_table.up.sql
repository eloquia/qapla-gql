CREATE SCHEMA core_qapla;

CREATE TABLE core_qapla.users (
  user_id     BIGSERIAL     NOT NULL,
  first_name  VARCHAR(64)   NOT NULL,
  last_name   VARCHAR(64)   NOT NULL,
  email       VARCHAR(128)  NOT NULL UNIQUE,
  created_at  timestamptz   NOT NULL DEFAULT NOW(),
  updated_at  timestamptz,
  PRIMARY KEY (user_id)
);

CREATE TABLE core_qapla.user_details (
  user_detail_id  BIGSERIAL     NOT NULL,
  user_id         BIGSERIAL     NOT NULL,
  middle_name     VARCHAR(64),
  goes_by         VARCHAR(64),
  gender          VARCHAR(64),
  ethnicity       VARCHAR(64),
  position        VARCHAR(64),
  institution     VARCHAR(64),
  PRIMARY KEY (user_detail_id),
  CONSTRAINT user_detail_fk FOREIGN KEY (user_id) REFERENCES core_qapla.users(user_id)
);

CREATE TABLE core_qapla.user_auths (
  user_auth_id        BIGSERIAL   NOT NULL,
  user_id             BIGSERIAL   NOT NULL,
  user_auth_password  VARCHAR(64) NOT NULL,
  created_at          timestamptz NOT NULL DEFAULT NOW(),
  updated_at          timestamptz,
  PRIMARY KEY (user_auth_id),
  CONSTRAINT user_auth_fk FOREIGN KEY (user_id) REFERENCES core_qapla.users(user_id)
);
