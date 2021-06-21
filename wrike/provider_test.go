package wrike

import (
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	os.Setenv("WRIKE_KEY", "bearer eyJ0dCI6InAiLCJhbGciOiJIUzI1NiIsInR2IjoiMSJ9.eyJkIjoie1wiYVwiOjQ2NTAxODYsXCJpXCI6Nzc0MzU1NyxcImNcIjo0NjI2MjkwLFwidVwiOjEwNjI2NjMwLFwiclwiOlwiVVNcIixcInNcIjpbXCJXXCIsXCJGXCIsXCJJXCIsXCJVXCIsXCJLXCIsXCJDXCIsXCJEXCIsXCJNXCIsXCJBXCIsXCJMXCIsXCJQXCJdLFwielwiOltdLFwidFwiOjB9IiwiaWF0IjoxNjIxMTc5NTkyfQ.56vbcUlIBctouj49OcOQoID0ehSmq4DveZHjKX3J2jY")
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"wrike": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		log.Println("[ERROR]: ", err)
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("WRIKE_KEY"); v == "" {
		t.Fatal("WRIKE_KEY must be set for acceptance tests")
	}
}
