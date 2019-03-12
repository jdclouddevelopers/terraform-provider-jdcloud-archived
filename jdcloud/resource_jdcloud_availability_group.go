package jdcloud

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/ag/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/ag/client"
)

func resourceJDCloudAvailabilityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceJDCloudAvailabilityGroupCreate,
		Read:   resourceJDCloudAvailabilityGroupRead,
		Delete: resourceJDCloudAvailabilityGroupDelete,

		Schema: map[string]*schema.Schema{
			"availability_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"az": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				MinItems: 1,
				ForceNew: true,
			},
			"instance_template_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ag_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "kvm",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceJDCloudAvailabilityGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*JDCloudConfig)

	req := apis.NewCreateAgRequest(config.Region)
	req.SetAzs(typeSetToStringArray(d.Get("az").(*schema.Set)))
	req.SetAgName(d.Get("availability_group_name").(string))
	req.SetInstanceTemplateId(d.Get("instance_template_id").(string))
	req.SetAgType(d.Get("ag_type").(string))
	if _, ok := d.GetOk("description"); ok {
		req.SetDescription(d.Get("description").(string))
	}

	agClient := client.NewAgClient(config.Credential)
	resp, err := agClient.CreateAg(req)
	if err != nil {
		return fmt.Errorf("[ERROR] create ag failed %s ", err.Error())
	}

	if resp.Error.Code != REQUEST_COMPLETED {
		return fmt.Errorf("[ERROR] create ag failed code:%d staus:%s message:%s", resp.Error.Code, resp.Error.Status, resp.Error.Message)
	}

	d.SetId(resp.Result.AgId)

	return nil
}

func resourceJDCloudAvailabilityGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*JDCloudConfig)
	req := apis.NewDeleteAgRequest(config.Region, d.Id())
	agClient := client.NewAgClient(config.Credential)

	resp, err := agClient.DeleteAg(req)

	if err != nil {
		return fmt.Errorf("[DEBUG]  delete ag failed %s", err.Error())
	}
	if resp.Error.Code != REQUEST_COMPLETED {
		return fmt.Errorf("[DEBUG] delete ag failed code:%d staus:%s message:%s", resp.Error.Code, resp.Error.Status, resp.Error.Message)
	}

	return nil
}

func resourceJDCloudAvailabilityGroupRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
