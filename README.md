# terraform-provider-powerdns

Simple provider to add A records through the PowerDNS REST API.

```
provider "powerdns" {
  api_url = "http://localhost:8081/servers/localhost/zones"
}

resource "powerdns_a_record" "test" {
  name = "rob.test"
  ip = "127.0.0.1"
}
```

## Install

```
go get github.com/tools/godep
go get github.com/roobert/terraform-provider-powerdns
cd $GOPATH/src/github.com/roobert/terraform-provider-powerdns
godep get
go build
go install
```
