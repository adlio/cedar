-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE projects (
  id serial NOT NULL PRIMARY KEY,
  tenant_id int NOT NULL,
  client_id int NOT NULL,
  parent_id int,
  name text NOT NULL DEFAULT '',
  slug text NOT NULL DEFAULT '',
  starts_on date,
  ends_on date,
  is_active boolean NOT NULL DEFAULT TRUE,
  is_billable boolean NOT NULL DEFAULT TRUE,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  FOREIGN KEY (tenant_id) REFERENCES tenants(id),
  FOREIGN KEY (client_id) REFERENCES clients(id),
  FOREIGN KEY (parent_id) REFERENCES projects(id),
  UNIQUE (tenant_id, client_id, parent_id, name)
);
CREATE INDEX ON projects (tenant_id, id);
CREATE INDEX ON projects (is_active);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE projects;
