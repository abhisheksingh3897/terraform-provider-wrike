package wrike

import (
	"terraform-provider-wrike/client"
	"terraform-provider-wrike/token"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"refresh_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("WRIKE_REFRESH_TOKEN", ""),
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("WRIKE_CLIENT_ID", ""),
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("WRIKE_CLIENT_SECRET", ""),
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
	token, err := token.GenerateToken(d.Get("client_id").(string), d.Get("client_secret").(string), d.Get("refresh_token").(string))
	if err != nil {
		return nil, err
	}
	return client.NewClient(token), nil

}
