resource "graylog_role" "foo" {
  name        = "foo"
  description = "Allows reading and writing all views and extended searches (built-in)"

  permissions = [
    "view:edit",
    "extendedsearch:use",
    "view:create",
    "extendedsearch:create",
    "view:read",
    "view:use"
  ]

  read_only = true
}
