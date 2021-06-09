package wrike

import (
	"terraform-provider-wrike/client"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"userid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"firstname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lastname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"accountid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Create: resourceCreateUser,
		Read:   resourceReadUser,
		Update: resourceUpdateUser,
		Delete: resourceDeleteUser,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateUser(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	err := apiClient.NewUser(d.Get("email").(string))
	if err != nil {
		return err
	}
	d.SetId(d.Get("email").(string))
	return nil
}

func resourceReadUser(d *schema.ResourceData, m interface{}) error {
	time.Sleep(60 * time.Second)
	apiClient := m.(*client.Client)
	UserId := d.Id()
	user, err := apiClient.GetUser(UserId)
	if err != nil {
		return err
	}
	d.Set("userid", user.ID)
	d.Set("firstname", user.FirstName)
	d.Set("lastname", user.LastName)
	d.Set("email", user.Profile[0].Email)
	d.Set("accountid", user.Profile[0].AccountID)
	return nil
}

func resourceUpdateUser(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDeleteUser(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
