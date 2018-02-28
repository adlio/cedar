-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE timesheet_entries (
  id serial NOT NULL PRIMARY KEY,
  tenant_id int NOT NULL,
  project_id int NOT NULL,
  role_id int NOT NULL,
  description text NOT NULL DEFAULT '',

);
CREATE INDEX ON timesheet_entries (tenant_id, id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

