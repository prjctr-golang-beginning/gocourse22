table "equipment" {
  schema = "public"
  column "equipment_id" {
    type = "int"
    null = false
    attrs = [auto_increment()]
  }
  column "name" {
    type = "varchar(100)"
    null = false
  }
  column "type" {
    type = "varchar(50)"
    null = false
  }
  column "quantity" {
    type = "int"
    null = false
  }
  column "status" {
    type = "varchar(50)"
    null = false
  }
  column "notes" {
    type = "text"
  }
  primary_key {
    columns = ["equipment_id"]
  }
}