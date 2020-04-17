package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/scalechamp/goss"
)

func ProviderFunc() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCALECHAMP_TOKEN", nil),
				Description: "Key used to authentication to the ScaleChamp API, retrived from project API settings",
			},
			"base_url": {
				Type:        schema.TypeString,
				Default:     "https://api.scalechamp.com",
				Optional:    true,
				Description: "Base URL to ScaleChamp API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"scalechamp_redis":      resourceInstance("redis"),
			"scalechamp_postgresql": resourceInstance("pg"),
			"scalechamp_mysql":      resourceInstance("mysql"),
			"scalechamp_keydb_pro":  resourceInstance("keydb-pro"),
			"scalechamp_keydb":      resourceInstance("keydb"),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return goss.NewClient(d.Get("base_url").(string), d.Get("token").(string)), nil
}
