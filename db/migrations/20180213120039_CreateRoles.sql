
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE roles (
  id serial NOT NULL,
  tenant_id int NOT NULL,
  name text NOT NULL DEFAULT '',
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  PRIMARY KEY(tenant_id, id),
  UNIQUE(tenant_id, name)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE roles;
