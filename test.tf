provider "powerdns" {
  api_url = "http://localhost:8081/servers/localhost/zones"
  api_key = "secret"
}
resource "powerdns_a_record" "test" {
  name = "rob.test"
  ip = "127.0.0.1"
}
