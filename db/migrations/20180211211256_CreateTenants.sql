-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tenants (
  id SERIAL NOT NULL PRIMARY KEY,
  subdomain text NOT NULL DEFAULT '' UNIQUE,
  company_name text NOT NULL DEFAULT '',
  time_zone text NOT NULL DEFAULT '',
  is_active boolean NOT NULL DEFAULT FALSE,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now()
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tenants;
