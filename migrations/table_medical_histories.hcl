table "medical_histories" {
  schema = "public"
  column "history_id" {
    type = "int"
    null = false
    attrs = [auto_increment()]
  }
  column "patient_id" {
    type = "int"
    null = false
  }
  column "disease" {
    type = "varchar(100)"
    null = false
  }
  column "treatment" {
    type = "text"
  }
  column "start_date" {
    type = "date"
  }
  column "end_date" {
    type = "date"
  }
  column "notes" {
    type = "text"
  }
  primary_key {
    columns = ["history_id"]
  }
  foreign_key "fk_patient_history" {
    columns     = ["patient_id"]
    ref_columns = ["patients.patient_id"]
  }
}