table "visits" {
  schema = "public"
  column "visit_id" {
    type = "int"
    null = false
    attrs = [auto_increment()]
  }
  column "patient_id" {
    type = "int"
    null = false
  }
  column "doctor_id" {
    type = "int"
    null = false
  }
  column "visit_date" {
    type = "timestamp with time zone"
    null = false
  }
  column "reason" {
    type = "text"
  }
  column "diagnosis" {
    type = "text"
  }
  column "notes" {
    type = "text"
  }
  primary_key {
    columns = ["visit_id"]
  }
  foreign_key "fk_patient" {
    columns     = ["patient_id"]
    ref_columns = ["patients.patient_id"]
  }
  foreign_key "fk_doctor" {
    columns     = ["doctor_id"]
    ref_columns = ["doctors.doctor_id"]
  }
}