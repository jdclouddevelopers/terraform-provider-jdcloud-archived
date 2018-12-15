package jdcloud

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type (
	JDCloudConfig struct {
		AccessKey  string
		SecretKey  string
		Region     string
		Credential *core.Credential
	}
)

var (
	regionCn = map[string]string{
		"cn-north-1": "华北-北京",
		"cn-south-1": "华南-广州",
		"cn-east-1":  "华东-宿迁",
		"cn-east-2":  "华东-上海",
	}
)

const (
	MAX_SECURITY_GROUP_COUNT = 5
	MAX_RECONNECT_COUNT      = 3

	REQUEST_COMPLETED  = 0
	RESOURCE_EXISTS    = 0
	RESOURCE_EMPTY     = 0
	RESOURCE_NOT_FOUND = 404
	REQUEST_INVALID    = 400

	MAX_DISK_COUNT          = 1
	DISK_AVAILABLE          = "available"
	DISK_DELETED            = "deleted"
	DISK_TIMEOUT            = 60
	DISK_ATTACHMENT_TIMEOUT = 60
	DISK_ATTACHED           = "attached"
	DISK_DETACHED           = "detached"

	MAX_EIP_COUNT = 1

	DISKTYPE_CLOUD = "cloud"
	MAX_VM_COUNT   = 1
	VM_TIMEOUT     = 600
	VM_PENDING     = "pending"
	VM_STARTING    = "starting"
	VM_RUNNING     = "running"
	VM_STOPPING    = "stopping"
	VM_STOPPED     = "stopped"
	VM_REBOOTING   = "rebooting"
	VM_REBUILDING  = "rebuilding"
	VM_RESIZING    = "resizing"
	VM_DELETING    = "deleting"
	VM_TERMINATED  = "terminated"
	VM_ERROR       = "error"
	VM_DELETED     = "deleted" //actual,there is no such state

	KEYPAIRS_PERM = 0600
	KEYPAIRS_PRIV = 0400

	RDS_TIMEOUT = 300
	RDS_READY   = "RUNNING"
	RDS_DELETED = ""
)

func initConfig(d *schema.ResourceData) (interface{}, error) {
	region := d.Get("region").(string)
	if _, ok := regionCn[region]; !ok {
		return nil, fmt.Errorf("Invalid region '%s'", region)
	}

	conf := &JDCloudConfig{
		AccessKey: d.Get("access_key").(string),
		SecretKey: d.Get("secret_key").(string),
		Region:    region,
		Credential: core.NewCredentials(
			d.Get("access_key").(string),
			d.Get("secret_key").(string),
		),
	}
	return conf, nil
}
