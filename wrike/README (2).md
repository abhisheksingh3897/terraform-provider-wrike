# Acceptance Testing

<https://www.terraform.io/docs/extend/testing/acceptance-tests/index.html>

#### Commands to run the acceptance test

1. go test -v <br/>
2. go test -v -cover (to get idea about how much percentage of your code is tested) <br />


### Different Testing Function

<strong>1. testAccCheckUserBasic </strong>

Checks the user exist exist or not in terraform state file.<br />

<strong>2. TestAccUser_basic </strong>

Creates the data source block and verifies that the returned resource attributes match.<br />

<strong>3. TestAccItem_Basic </strong>

Creates the resource block and verifies that the returned resource attributes match. <br />