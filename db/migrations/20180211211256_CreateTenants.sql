-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tenants (
  id SERIAL NOT NULL PRIMARY KEY,
  subdomain text NOT NULL UNIQUE,
  company_name text NOT NULL,
  time_zone text NOT NULL,
  is_active boolean NOT NULL DEFAULT FALSE,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now()
);
CREATE INDEX ON tenants (subdomain);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tenants;
