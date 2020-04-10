provider "influxdbv2-onboarding" {
  url = "http://localhost:9999"
}

resource "influxdbv2-onboarding_setup" "setup" {
  username = "test"
  password = "test1234"
  bucket = "test-bucket"
  org = "test-org"
  retention_period = 4
}

output "token" {
  value = influxdbv2-onboarding_setup.setup.token
}
output "user_id" {
  value = influxdbv2-onboarding_setup.setup.user_id
}
output "org_id" {
  value = influxdbv2-onboarding_setup.setup.org_id
}
output "bucket_id" {
  value = influxdbv2-onboarding_setup.setup.bucket_id
}
output "auth_id" {
  value = influxdbv2-onboarding_setup.setup.auth_id
}