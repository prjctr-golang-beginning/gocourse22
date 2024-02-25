table "business_rule_check" {
  schema = schema.public
  column "id" {
    type    = uuid
    default = sql("gen_random_uuid()")
    null    = false
  }

  column "business_rule_id" {
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

  column "type" {
    type = enum.check_type
    null = false
  }

  column "type_config" {
    type = json
    null = false
  }

  column "last_run_date" {
    type = timestamp
    null = true
  }

  column "last_run_result" {
    type = enum.check_result
    null = true
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

  primary_key {
    columns = [column.id]
  }

  foreign_key "business_rule_id" {
    columns     = [column.business_rule_id]
    ref_columns = [table.business_rule.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }

  comment = "Business rule check table"
}