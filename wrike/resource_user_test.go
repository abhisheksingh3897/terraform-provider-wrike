package wrike

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccItem_Basic(t *testing.T) {
	os.Setenv("TF_ACC", "1")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("wrike_user.user1", "email", "abhisheksinghiiitp@gmail.com"),
				),
			},
		},
	})
}

func testAccCheckItemBasic() string {
	return fmt.Sprintf(`
resource "wrike_user" "user1" {    
    email= "abhisheksinghiiitp@gmail.com"
}
`)
}

func TestAccWrikeUserResource_update(t *testing.T) {
	os.Setenv("TF_ACC", "1")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWrikeUserResourceUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("wrike_user.user1", "email", "abhisheksinghiiitp@gmail.com"),
					resource.TestCheckResourceAttr("wrike_user.user1", "role", "User"),
					resource.TestCheckResourceAttr("wrike_user.user1", "external", "false"),
				),
			},
			{
				Config: testAccCheckWrikeUserResourceUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("wrike_user.user1", "email", "abhisheksinghiiitp@gmail.com"),
					resource.TestCheckResourceAttr("wrike_user.user1", "role", "Collaborator"),
					resource.TestCheckResourceAttr("wrike_user.user1", "external", "true"),
				),
			},
		},
	})
}

func testAccCheckWrikeUserResourceUpdatePre() string {
	return fmt.Sprintf(`
	resource "wrike_user" "user1" {    
		email= "abhisheksinghiiitp@gmail.com"
		role= "User"
		external="false"
	}`)
}

func testAccCheckWrikeUserResourceUpdatePost() string {
	return fmt.Sprintf(`
	resource "wrike_user" "user1" {    
		email= "abhisheksinghiiitp@gmail.com"
		role= "Collaborator"
		external="true"
	}`)
}
