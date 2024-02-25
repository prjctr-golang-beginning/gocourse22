table "scan_raw_result" {
  schema = schema.public
  column "id" {
    type    = uuid
    default = sql("gen_random_uuid()")
  }

  column "job_id" {
    type = uuid
    null = false
  }

  column "business_rule_id" {
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

  column "run_status" {
    type = varchar(255)
    null = false
  }

  column "state" {
    type    = enum.scan_raw_result_state
    null    = false
  }

  column "parsed_at" {
    type = timestamp
    null = true
  }

  column "return_code" {
    type = int
    null = false
  }

  column "scan_result" {
    type = json
    null = true
  }

  column "std_out" {
    type = text
    null = true
  }

  column "std_err" {
    type = text
    null = true
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

  foreign_key "business_rule_id" {
    columns     = [column.business_rule_id]
    ref_columns = [table.business_rule.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }

  comment = "Scan result table"
}