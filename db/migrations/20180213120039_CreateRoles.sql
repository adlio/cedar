-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE roles (
  id serial NOT NULL PRIMARY KEY,
  tenant_id int NOT NULL,
  name text NOT NULL,
  is_active boolean NOT NULL DEFAULT TRUE,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  FOREIGN KEY (tenant_id) REFERENCES tenants(id),
  UNIQUE(tenant_id, name)
);
CREATE INDEX ON roles (tenant_id, id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE roles;
