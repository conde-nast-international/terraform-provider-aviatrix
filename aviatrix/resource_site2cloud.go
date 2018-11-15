package aviatrix

import (
	"fmt"
	"log"

	"github.com/AviatrixSystems/go-aviatrix/goaviatrix"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAviatrixSite2Cloud() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviatrixSite2CloudCreate,
		Read:   resourceAviatrixSite2CloudRead,
		Update: resourceAviatrixSite2CloudUpdate,
		Delete: resourceAviatrixSite2CloudDelete,

		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"connection_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_gateway_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"connection_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tunnel_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"primary_cloud_gateway_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backup_gateway_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pre_shared_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gateway_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_subnet_cidr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"local_subnet_cidr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha_enabled": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_remote_subnet_cidr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_remote_gateway_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_remote_gateway_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_pre_shared_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_route_encryption": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAviatrixSite2CloudCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)
	s2c := &goaviatrix.Site2Cloud{
		GwName:             d.Get("primary_cloud_gateway_name").(string),
		BackupGwName:       d.Get("backup_gateway_name").(string),
		VpcID:              d.Get("vpc_id").(string),
		TunnelName:         d.Get("connection_name").(string),
		ConnType:           d.Get("connection_type").(string),
		TunnelType:         d.Get("tunnel_type").(string),
		RemoteGwType:       d.Get("remote_gateway_type").(string),
		RemoteGwIP:         d.Get("remote_gateway_ip").(string),
		RemoteGwIP2:        d.Get("backup_remote_gateway_ip").(string),
		PreSharedKey:       d.Get("pre_shared_key").(string),
		BackupPreSharedKey: d.Get("backup_pre_shared_key").(string),
		RemoteSubnet:       d.Get("remote_subnet_cidr").(string),
		LocalSubnet:        d.Get("local_subnet_cidr").(string),
		HAEnabled:          d.Get("ha_enabled").(string),
	}

	log.Printf("[INFO] Creating Aviatrix Site2Cloud: %#v", s2c)

	err := client.CreateSite2Cloud(s2c)
	if err != nil {
		return fmt.Errorf("failed Site2Cloud create: %s", err)
	}
	d.SetId(s2c.TunnelName + s2c.VpcID)
	return resourceAviatrixSite2CloudRead(d, meta)
}

func resourceAviatrixSite2CloudRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)
	site2cloud := &goaviatrix.Site2Cloud{
		TunnelName: d.Get("connection_name").(string),
		VpcID:      d.Get("vpc_id").(string),
	}
	s2c, err := client.GetSite2Cloud(site2cloud)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("couldn't find Aviatrix Site2Cloud: %s, %#v", err, s2c)
	}
	if s2c != nil {
		d.Set("vpc_id", s2c.VpcID)
		d.Set("remote_gateway_type", s2c.RemoteGwType)
		d.Set("tunnel_type", s2c.TunnelType)
		d.Set("remote_gateway_ip", s2c.RemoteGwIP)
		d.Set("remote_subnet_cidr", s2c.RemoteSubnet)
		if d.Get("local_subnet_cidr") != "" {
			d.Set("local_subnet_cidr", s2c.LocalSubnet)
		}
		if connectionType := d.Get("connection_type").(string); connectionType == "" {
			//force default setting and save to .tfstate file
			d.Set("connection_type", "unmapped")
		}
	}
	log.Printf("[TRACE] Reading Aviatrix Site2Cloud %s: %#v", d.Get("connection_name").(string), site2cloud)
	log.Printf("[TRACE] Reading Aviatrix Site2Cloud connection_type: [%s]", d.Get("connection_type").(string))
	d.SetId(site2cloud.TunnelName + site2cloud.VpcID)
	return nil
}

func resourceAviatrixSite2CloudUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)
	site2cloud := &goaviatrix.Site2Cloud{
		GwName:     d.Get("primary_cloud_gateway_name").(string),
		VpcID:      d.Get("vpc_id").(string),
		TunnelName: d.Get("connection_name").(string),
	}
	d.Partial(true)
	log.Printf("[INFO] Updating Aviatrix Site2Cloud: %#v", site2cloud)
	if ok := d.HasChange("remote_subnet_cidr"); ok {
		site2cloud.RemoteSubnet = d.Get("remote_subnet_cidr").(string)
		err := client.UpdateSite2Cloud(site2cloud)
		if err != nil {
			return fmt.Errorf("failed to update Site2Cloud remote_subnet_cidr: %s", err)
		}
		d.SetPartial("remote_subnet_cidr")
	}
	if ok := d.HasChange("local_subnet_cidr"); ok {
		site2cloud.LocalSubnet = d.Get("local_subnet_cidr").(string)
		err := client.UpdateSite2Cloud(site2cloud)
		if err != nil {
			return fmt.Errorf("failed to update Site2Cloud local_subnet_cidr: %s", err)
		}
		d.SetPartial("local_subnet_cidr")
	}
	d.Partial(false)
	d.SetId(site2cloud.TunnelName + site2cloud.VpcID)
	return resourceAviatrixSite2CloudRead(d, meta)
}

func resourceAviatrixSite2CloudDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)
	s2c := &goaviatrix.Site2Cloud{
		VpcID:      d.Get("vpc_id").(string),
		TunnelName: d.Get("connection_name").(string),
	}

	log.Printf("[INFO] Deleting Aviatrix s2c: %#v", s2c)

	err := client.DeleteSite2Cloud(s2c)
	if err != nil {
		return fmt.Errorf("failed to delete Aviatrix Site2Cloud: %s", err)
	}
	return nil
}
