table "schedules" {
  schema = "public"
  column "schedule_id" {
    type = "int"
    null = false
    attrs = [auto_increment()]
  }
  column "doctor_id" {
    type = "int"
    null = false
  }
  column "work_day" {
    type = "date"
    null = false
  }
  column "start_time" {
    type = "time"
    null = false
  }
  column "end_time" {
    type = "time"
    null = false
  }
  column "notes" {
    type = "text"
  }
  primary_key {
    columns = ["schedule_id"]
  }
  foreign_key "fk_doctor_schedule" {
    columns     = ["doctor_id"]
    ref_columns = ["doctors.doctor_id"]
  }
}