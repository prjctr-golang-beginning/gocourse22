table "jobs" {
  schema = schema.public
  column "id" {
    type    = uuid
    default = sql("gen_random_uuid()")
  }

  column "name" {
    type = varchar(255)
    null = false
  }

  column "status" {
    type    = enum.job_status
    null    = false
    default = "new"
  }

  column "worker_id" {
    type    = int
    null    = true
  }

  column "started_at" {
    type = timestamp
    null = true
    comment = "seconds"
  }

  column "process_took" {
    type = double_precision
    null = true
  }

  column "schedule_type" {
    type = varchar(255)
    null = false
  }

  column "config" {
    type = json
    null = false
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

  comment = "Jobs table"
}