package main

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/scalechamp/goss"
	"time"
)

func resourceInstance(kind string) *schema.Resource {
	i := &instanceResource{kind}
	s := map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the instanceResource",
		},
		"plan": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the plan, valid options are: lemur, tiger, bunny, rabbit, panda, ape, hippo, lion",
		},
		"cloud": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the region you want to create your instanceResource i",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Name of the region you want to create your instanceResource i",
		},
		"whitelist": {
			Type:        schema.TypeSet,
			Elem:        &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "Name of the region you want to create your instanceResource i",
		},
		"region": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "API key for the CloudAMQP instanceResource",
		},
		"master_host": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "API key for the CloudAMQP instanceResource",
		},
		"replica_host": {
			Type:        schema.TypeString,
			Computed:    true,
			Optional:    true,
			Description: "API key for the CloudAMQP instanceResource",
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "API key for the CloudAMQP instanceResource",
		},
	}
	if kind == "keydb-pro" {
		s["license_key"] = &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the instanceResource",
		}
	}
	if kind == "keydb" || kind == "keydb-pro" || kind == "redis" {
		s["eviction_policy"] = &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the instanceResource",
		}
	}
	return &schema.Resource{
		Create: i.resourceCreate,
		Read:   i.resourceRead,
		Update: i.resourceUpdate,
		Delete: i.resourceDelete,
		Schema: s,
	}
}

type instanceResource struct {
	kind string
}

func (r *instanceResource) resourceCreate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*goss.Client)
	plan, err := api.Plans.Find(&goss.PlanFindRequest{
		Cloud:  d.Get("cloud").(string),
		Region: d.Get("region").(string),
		Name:   d.Get("plan").(string),
		Kind:   r.kind,
	})
	if err != nil {
		return err
	}
	data, err := api.Instances.Create(&goss.InstanceCreateRequest{
		Name:           d.Get("name").(string),
		Whitelist:      d.Get("whitelist").([]string),
		PlanID:         plan.ID,
	})
	if err != nil {
		return err
	}
	d.SetId(data.ID)

	for i := 0; i < 15; i += 1 {
		data, err = api.Instances.Get(data.ID)
		if err != nil {
			return err
		}
		if data.State == "running" {
			break
		}
		if data.State == "failed" {
			return errors.New("failed instanceResource")
		}
		time.Sleep(30 * time.Second)
	}

	return r.resourceRead(d, meta)

}

func (r *instanceResource) resourceUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*goss.Client)
	instanceUpdateRequest := new(goss.InstanceUpdateRequest)
	instanceUpdateRequest.ID = d.Id()

	if d.HasChange("name") {
		name := d.Get("name").(string)
		instanceUpdateRequest.Name = &name
	}

	if d.HasChange("cloud") || d.HasChange("region") || d.HasChange("plan") {
		plan, err := api.Plans.Find(&goss.PlanFindRequest{
			Cloud:  d.Get("cloud").(string),
			Region: d.Get("region").(string),
			Name:   d.Get("plan").(string),
			Kind:   r.kind,
		})
		if err != nil {
			return err
		}
		instanceUpdateRequest.PlanID = plan.ID
	}

	if d.HasChange("password") {
		password := d.Get("password").(string)
		instanceUpdateRequest.Password = password
	}

	if d.HasChange("enabled") {
		enabled := d.Get("enabled").(bool)
		instanceUpdateRequest.Enabled = &enabled
	}

	if d.HasChange("whitelist") {
		whitelist := d.Get("whitelist").([]string)
		instanceUpdateRequest.Whitelist = &whitelist
	}

	_, err := api.Instances.Update(instanceUpdateRequest)
	if err != nil {
		return err
	}
	for i := 0; i < 50; i += 1 {
		time.Sleep(5 * time.Second)
		instance, err := api.Instances.Get(d.Id())
		if err != nil {
			return err
		}
		if instance.State == "running" {
			break
		}
		if instance.State == "failed" {
			return errors.New("failed instanceResource")
		}
	}
	return r.resourceRead(d, meta)
}

func (r *instanceResource) resourceRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*goss.Client)
	instance, err := api.Instances.Get(d.Id())
	if err != nil {
		return err
	}
	d.Set("replica_host", instance.ConnectionInfo.ReplicaHost)
	d.Set("master_host", instance.ConnectionInfo.MasterHost)
	d.Set("password", instance.Password)
	return nil
}

func (r *instanceResource) resourceDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*goss.Client)
	return api.Instances.Delete(d.Id())
}
