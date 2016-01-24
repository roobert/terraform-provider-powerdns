package powerdns

import (
	"github.com/hashicorp/terraform/helper/schema"
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
				Type:     schema.TypeString,
				Optional: true,
				Default:  86400,
			},
		},
	}
}

func resourceARecordCreate(d *schema.ResourceData, m interface{}) error {
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
