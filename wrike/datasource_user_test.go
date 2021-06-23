package wrike

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUser_Basic(t *testing.T) {
	os.Setenv("TF_ACC", "1")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUserBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.wrike_user.user", "email", "abhishek.s@clevertap.com"),
					resource.TestCheckResourceAttr("data.wrike_user.user", "firstname", "Abhishek"),
					resource.TestCheckResourceAttr("data.wrike_user.user", "lastname", "singh"),
				),
			},
		},
	})
}

func testAccCheckUserBasic() string {
	return fmt.Sprintf(`
		data "wrike_user" "user" {    
    		email= "abhishek.s@clevertap.com"
		}
	`)
}
