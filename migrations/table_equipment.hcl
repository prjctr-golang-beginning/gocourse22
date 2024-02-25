table "jobs_run_history" {
  schema = schema.public
  column "id" {
    type    = uuid
    default = sql("gen_random_uuid()")
  }

  column "job_id" {
    type    = uuid
    null    = false
  }

  column "status" {
    type    = enum.job_run_status
    null    = false
  }

  column "created_at" {
    type = timestamp
    null = false
    default = sql("now()")
  }

  column "details" {
    type = text
    null = true
  }

  primary_key {
    columns = [column.id]
  }

  comment = "Jobs run history table"
}