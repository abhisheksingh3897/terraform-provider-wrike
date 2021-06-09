package wrike

import (
	"terraform-provider-wrike/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceUser() *schema.Resource {
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
				Required: true,
			},
		},
		Read: datasourceReadUser,
	}
}

func datasourceReadUser(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	UserId := d.Get("email").(string)
	user, err := apiClient.GetUser(UserId)
	if err != nil {
		return err
	}
	d.Set("userid", user.ID)
	d.Set("firstname", user.FirstName)
	d.Set("lastname", user.LastName)
	d.Set("email", user.Profile[0].Email)
	d.Set("accountid", user.Profile[0].AccountID)
	d.SetId(UserId)
	return nil
}
