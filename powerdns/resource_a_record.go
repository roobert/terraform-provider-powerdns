package powerdns

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/parnurzeal/gorequest"
  "encoding/json"
	"log"
)

func resourceARecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceARecordCreate,
		Read:   resourceARecordRead,
		Update: resourceARecordUpdate,
		Delete: resourceARecordDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
		},
	}
}

type RRSets struct {
	RRSets []RRSet
}

type RRSet struct {
	Type       string `json: "type"`
	name       string
	changetype string
	records    []Record
}

type Record struct {
	Type     string `json: "type"`
	content  string
	disabled bool
	name     string
	ttl      int
}

func resourceARecordCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)

	name := d.Get("name").(string)

	record := Record{
		content:  d.Get("ip").(string),
		disabled: false,
		name:     name,
		ttl:      d.Get("ttl").(int),
	}

  // FIXME: add error checking
  json.Unmarshal([]byte(`{"type": "A"}`), &record)

	records := []Record{record}

	rrset := RRSet{
		name:       name,
		changetype: "REPLACE",
		records:    records,
	}

  // FIXME: add error checking
  json.Unmarshal([]byte(`{"type": "A"}`), &rrset)

	rrsets := []RRSet{rrset}
	data := RRSets{RRSets: rrsets}

	log.Printf("PowerDNS API URL: %s\n", config.APIUrl)
	log.Printf("PowerDNS API post data: %+v\n", data)

	request := gorequest.New()
	resp, body, err := request.Post(config.APIUrl).
		Set("X-API-Key", config.APIKey).
		Send(data).
		End()

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	if resp.Status != "200" {
		return fmt.Errorf("Error: (%s): %s", resp.Status, body)
	}

	log.Printf("PowerDNS API body: %s\n", body)
	log.Printf("PowerDNS API response: %+v\n", resp)

	return nil
}

func resourceARecordRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceARecordUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceARecordDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
