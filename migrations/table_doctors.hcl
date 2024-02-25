table "business_rule" {
  schema = schema.public
  column "id" {
    type    = uuid
    default = sql("gen_random_uuid()")
    null    = false
  }

  column "validation_suite_id" {
    type = uuid
    null = false
  }

  column "name" {
    type = varchar(255)
    null = false
  }

  column "description" {
    type = varchar(512)
    null = true
  }

  column "launch_schedule" {
    type = varchar(64)
    null = true
  }

  column "launch_status" {
    type    = enum.launch_status
    null    = false
    default = "standby"
  }

  column "datasource_id" {
    type = uuid
    null = false
  }

  column "status" {
    type    = enum.status
    null    = false
    default = "disabled"
  }

  column "created_by" {
    type = varchar(26)
    null = false
  }

  column "created_at" {
    type = timestamp
    null = false
  }

  column "updated_at" {
    type = timestamp
    null = true
  }

  column "category" {
    type = enum.rule_category
    null = true
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "validation_suite_id" {
    columns     = [column.validation_suite_id]
    ref_columns = [table.validation_suite.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }

  foreign_key "datasource_id" {
    columns     = [column.datasource_id]
    ref_columns = [table.datasource.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }

  comment = "Business rule table"
}