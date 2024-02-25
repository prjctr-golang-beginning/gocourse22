schema "public" {
  comment = "standard public schema"
}

enum "status" {
  schema = schema.public
  values = ["enabled", "disabled"]
}
