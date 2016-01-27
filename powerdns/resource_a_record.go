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
		Update: resourceARecordCreate,
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
	RRSets []RRSet `json:"rrsets"`
}

type RRSet struct {
	Type       string   `json:"type"`
	Name       string   `json:"name"`
	Changetype string   `json:"changetype"`
	Records    []Record `json:"records"`
}

type Record struct {
	Type     string `json:"type"`
	Content  string `json:"content"`
	Disabled bool   `json:"disabled"`
	Name     string `json:"name"`
	Ttl      int    `json:"ttl"`
}

func resourceARecordCreate(d *schema.ResourceData, m interface{}) error {

	// FIXME: at the moment CREATE and UPDATE are the same method because
	//        powerdns API expects changetype to be REPLACE when creating records.
	//
	//        it would make sense for the CREATE method to first check if a resource
	//        exists (using READ method?) and then call the UPDATE method.

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
		Changetype: "REPLACE",
		Records:    records,
		Type:       "A",
	}

	rrsets := []RRSet{rrset}
	data := RRSets{RRSets: rrsets}

	log.Printf("PowerDNS API URL: %s\n", config.APIUrl)
	log.Printf("PowerDNS API post data: %+v\n", data)

	url := config.APIUrl

	request := gorequest.New()
	resp, body, err := request.Patch(url).
		Set("X-API-Key", config.APIKey).
		Send(data).
		End()

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error: (%s): %s", resp.Status, body)
	}

	log.Printf("PowerDNS API body: %s\n", body)
	log.Printf("PowerDNS API response: %+v\n", resp)

	d.SetId("powerdns_a_record-" + url)

	return nil
}

func resourceARecordRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceARecordDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)

	name := d.Get("name").(string)

	records := []Record{}

	rrset := RRSet{
		Name:       name,
		Changetype: "DELETE",
		Records:    records,
		Type:       "A",
	}

	rrsets := []RRSet{rrset}
	data := RRSets{RRSets: rrsets}

	log.Printf("PowerDNS API URL: %s\n", config.APIUrl)
	log.Printf("PowerDNS API post data: %+v\n", data)

	url := config.APIUrl

	request := gorequest.New()
	resp, body, err := request.Patch(url).
		Set("X-API-Key", config.APIKey).
		Send(data).
		End()

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error: (%s): %s", resp.Status, body)
	}

	log.Printf("PowerDNS API body: %s\n", body)
	log.Printf("PowerDNS API response: %+v\n", resp)

	return nil
}
