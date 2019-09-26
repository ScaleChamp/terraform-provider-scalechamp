package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/scalablespace/goss"
)

func resourceInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: func(data *schema.ResourceData, i interface{}) error {
			return nil
		},
		Delete: resourceDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the instance",
			},
			"plan": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the plan, valid options are: lemur, tiger, bunny, rabbit, panda, ape, hippo, lion",
			},
			"dc": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the region you want to create your instance in",
			},
			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Sensitive:   true,
				Description: "API key for the CloudAMQP instance",
			},
		},
	}
}

func resourceCreate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*goss.API)
	keys := []string{"name", "plan", "dc"}
	params := make(map[string]interface{})
	for _, k := range keys {
		if v := d.Get(k); v != nil {
			params[k] = v
		}
	}
	data, err := api.CreateInstance(params)
	if err != nil {
		return err
	}
	d.SetId(data["id"].(string))
	for k, v := range data {
		d.Set(k, v)
	}
	return nil

}

func resourceRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*goss.API)
	data, err := api.ReadInstance(d.Id())
	if err != nil {
		return err
	}
	for k, v := range data {
		d.Set(k, v)
	}
	return nil
}

func resourceDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*goss.API)
	return api.DeleteInstance(d.Id())
}
