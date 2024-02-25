table "validation_suite" {
  schema = schema.public
  column "id" {
    type    = uuid
    default = sql("gen_random_uuid()")
  }

  column "account_id" {
    type = varchar(26)
    null = false
    default = ""
  }

  column "name" {
    type = varchar(255)
    null = false
  }

  column "description" {
    type = varchar(512)
    null = true
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

  primary_key {
    columns = [column.id]
  }

  comment = "Validation suite table"
}