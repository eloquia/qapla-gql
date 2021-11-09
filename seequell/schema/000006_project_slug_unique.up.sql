ALTER TABLE core_qapla.projects
ADD CONSTRAINT projects_unique_slug UNIQUE (slug);
