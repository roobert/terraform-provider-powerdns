provider "powerdns" {}

resource "a_record" "test" {
  name = "rob.test"
  ip = "127.0.0.1"
}
