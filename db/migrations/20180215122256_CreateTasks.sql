-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tasks (
  id serial NOT NULL PRIMARY KEY,
  tenant_id int NOT NULL,
  project_id int NOT NULL,
  parent_id int,
  name text NOT NULL,
  slug text NOT NULL,
  is_billable boolean NOT NULL DEFAULT TRUE,
  percent_complete NUMERIC(5,2) NOT NULL DEFAULT 0.0,
  estimate_in_hours NUMERIC(6,2) NOT NULL DEFAULT 0.0,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  FOREIGN KEY (parent_id) REFERENCES tasks(id),
  FOREIGN KEY (tenant_id) REFERENCES tenants(id),
  FOREIGN KEY (project_id) REFERENCES projects(id),
  UNIQUE (tenant_id, project_id, parent_id, name)
);
CREATE INDEX ON tasks (tenant_id, id);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tasks;
