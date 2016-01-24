package powerdns

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/parnurzeal/gorequest"
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
	Name       string
	Type       string
	Changetype string
	Records    []Record
}

type Record struct {
	Content  string
	Disabled bool
	Name     string
	Ttl      int
	Type     string
}

func resourceARecordCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)

	name := d.Get("name").(string)

	record := Record{
		Content:  d.Get("ip").(string),
		Disabled: false,
		Name:     name,
		Ttl:      d.Get("ttl").(int),
		Type:     "A",
	}

	records := []Record{record}

	rrset := RRSet{
		Name:       name,
		Type:       "A",
		Changetype: "REPLACE",
		Records:    records,
	}

	rrsets := []RRSet{rrset}
	data := RRSets{RRSets: rrsets}

	log.Printf("PowerDNS API URL: %s\n", config.APIUrl)
	log.Printf("PowerDNS API post data: %+v\n", data)

	request := gorequest.New()
	resp, body, err := request.Post(config.APIUrl).
		Set("X-API-Key", "testing_dns").
		Send(data).
		End()

	if err != nil {
		return fmt.Errorf("%s", err)
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
