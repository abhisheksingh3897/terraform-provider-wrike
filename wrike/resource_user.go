package wrike

import (
	"fmt"
	"strings"
	"terraform-provider-wrike/client"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

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
				Required: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
		Create: resourceCreateUser,
		Read:   resourceReadUser,
		Update: resourceUpdateUser,
		Delete: resourceDeleteUser,
		Importer: &schema.ResourceImporter{
			State: resourceUserImporter,
		},
	}
}

func resourceCreateUser(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	var err error
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		if err = apiClient.NewUser(d.Get("email").(string)); err != nil {
			if apiClient.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if retryErr != nil {
		time.Sleep(2 * time.Second)
		return retryErr
	}
	if err != nil {
		return err
	}
	d.SetId(d.Get("email").(string))
	return nil
}

func resourceReadUser(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	UserId := d.Id()

	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		user, err := apiClient.GetUser(UserId)
		if err != nil {
			if apiClient.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		d.Set("userid", user.ID)
		d.Set("firstname", user.FirstName)
		d.Set("lastname", user.LastName)
		d.Set("email", user.Profile[0].Email)
		d.Set("accountid", user.Profile[0].AccountID)
		d.Set("role", user.Profile[0].Role)
		d.Set("external", user.Profile[0].External)
		return nil
	})
	if retryErr != nil {
		if strings.Contains(retryErr.Error(), "user not found") == true {
			d.SetId("")
			return nil
		}
		return retryErr
	}
	return nil
}

func resourceUpdateUser(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("email") {
		return fmt.Errorf("User is not allowed to change email")
	}
	apiClient := m.(*client.Client)
	UserId := d.Id()
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		err := apiClient.UpdateUser(UserId, d.Get("accountid").(string), d.Get("role").(string), d.Get("external").(bool))
		if err != nil {
			if apiClient.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if retryErr != nil {
		if strings.Contains(retryErr.Error(), "user not found") == true {
			d.SetId("")
			return nil
		}
		return retryErr
	}
	return nil
}

func resourceDeleteUser(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func resourceUserImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	apiClient := m.(*client.Client)
	UserId := d.Id()
	user, err := apiClient.GetUser(UserId)
	if err != nil {
		return nil, err
	}
	d.Set("userid", user.ID)
	d.Set("firstname", user.FirstName)
	d.Set("lastname", user.LastName)
	d.Set("email", user.Profile[0].Email)
	d.Set("accountid", user.Profile[0].AccountID)
	d.Set("role", user.Profile[0].Role)
	d.Set("external", user.Profile[0].External)
	return []*schema.ResourceData{d}, nil
}
