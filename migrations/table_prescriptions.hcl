table "scan_result" {
  schema = schema.public
  column "id" {
    type    = uuid
    default = sql("gen_random_uuid()")
  }

  column "job_id" {
    type = uuid
    null = false
  }

  column "check_id" {
    type = uuid
    null = false
  }

  column "scan_start_time" {
    type = timestamp
    null = false
  }

  column "scan_end_time" {
    type = timestamp
    null = false
  }

  column "status" {
    type = varchar(255)
    null = false
  }

  column "has_errors" {
    type = boolean
    null = false
    default = false
  }

  column "has_warnings" {
    type = boolean
    null = false
    default = false
  }

  column "has_failures" {
    type = boolean
    null = false
    default = false
  }

  column "message" {
    type = text
    null = true
  }

  column "created_at" {
    type = timestamp
    null = false
    default = sql("now()")
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "job_id" {
    columns     = [column.job_id]
    ref_columns = [table.jobs.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }

  foreign_key "check_id" {
    columns     = [column.check_id]
    ref_columns = [table.business_rule_check.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }

  comment = "Scan raw result table"
}