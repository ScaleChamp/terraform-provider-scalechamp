package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/scalablespace/goss"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apikey": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCALABLESPACE_APIKEY", nil),
				Description: "Key used to authentication to the CloudAMQP Customer API",
			},
			"baseurl": {
				Type:        schema.TypeString,
				Default:     "https://api.scalablespace.net",
				Optional:    true,
				Description: "Base URL to CloudAMQP Customer website",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"scalablespace_instance": resourceInstance(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return goss.New(d.Get("baseurl").(string), d.Get("apikey").(string)), nil
}
