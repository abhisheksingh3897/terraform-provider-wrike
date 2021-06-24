package wrike

import (
	"terraform-provider-wrike/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("WRIKE_TOKEN", ""),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"wrike_user": datasourceUser(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"wrike_user": resourceUser(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	token := d.Get("token").(string)
	return client.NewClient(token), nil

}
