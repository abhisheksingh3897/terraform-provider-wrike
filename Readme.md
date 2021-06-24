This Terraform provider enables create, read, update and import operations for wrike users.

## Requirements

* [Go](https://golang.org/doc/install) >= 1.16 (To build the provider plugin)<br>
* [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x <br/>
* Application: [Wrike](https://www.wrike.com/)
* [Wrike API Documentation](https://developers.wrike.com/api/v4/users/)


## Application Account

### Setup<a id="setup"></a>
1. Create an wrike account at https://www.wrike.com/<br>

### API Authentication
1. To authenticate API we need a pair of credentials: Client ID and Client secret.<br>
2. To access client id and client secret, you need create an `wrike App` from Apps and Integration<br>
3. A pair of credentials: Client ID and Client secret will be generated and also write one Redirect URI (eg https://www.wrike.com/).<br>
4. With the help of pair of credential, you can generate `Permanent Access Token`.<br>


## Building The Provider
1. Clone the repository and add the dependencies and create a vendor directory that contains all dependencies. For this run the following commands: <br>
```
cd terraform-provider-wrike
go mod init terraform-provider-wrike
go mod tidy
go mod vendor
```


## Managing terraform plugins
*For Windows:*
1. Run the following command to create a vendor subdirectory (`%APPDATA%/terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${OS_ARCH}`) which will comprise of all terraform plugins. <br> 
Command: 
```bash
mkdir -p %APPDATA%/terraform.d/plugins/wrike.com\user\wrike\4.0.0\windows_amd64
```
2. Run `go build -o terraform-provider-wrike.exe` to generate the binary in present working directory. <br>
3. Run this command to move this binary file to appropriate location.
 ```
 move terraform-provider-expensify.exe %APPDATA%\terraform.d\plugins\wrike.com\user\wrike\4.0.0\windows_amd64
 ``` 
 <p align="center">[OR]</p>
 
3. Manually move the file from current directory to destination directory (`%APPDATA%\terraform.d\plugins\wrike.com\user\wrike\4.0.0\windows_amd64`).<br>


## Working with terraform

### Application Credential Integration in terraform
1. Add `terraform` block and `rovider` block as shown in [example usage](#example-usage).
2. Get a pair of credentials: Client ID and Client secret. By help of this, generate access token.
3. Assign the above Access Token to the repective field in the `provider` block.

### Basic Terraform Commands
1. `terraform init` - To initialize a working directory containing Terraform configuration files.
2. `terraform plan` - To create an execution plan. Displays the changes to be done.
3. `terraform apply` - To execute the actions proposed in a Terraform plan. Apply the chages.

#### Create User
1. Add the user `email` in the respective field as shown in the [example usage](#example-usage)
2. Run the basic terraform commands.<br>
3. On successful execution, sends an account setup mail to user.<br>

#### Read the User Data
Add `data` and `output` blocks as shown in the [example usage](#example-usage) and run the basic terraform commands.

### Update the user
1. Update the data of the user in the `resource` block as show in [example usage](#example-usage) and run the basic terraform commands to update user. 
2. We can only update the role of user i.e. `Regular user, External user and Collaborator`.
3. To make user as a Collaborator, external should be true.

#### Import a User Data
1. Write manually a `resource` configuration block for the user as shown in [example usage](#example-usage). Imported user will be mapped to this block.
2. Run the command `terraform import wrike_user.user “[EMAIL_ID]”`
3. Run `terraform plan`, if output show `0 to addd, 0 to change and 0 to destroy` user import is successful, otherwise recheck the employee data in resource block with employee data in the policy in Expensify Website.

## Example Usage<a id="example-usage"></a>

```terraform
terraform {
	required_providers {
	  wrike = {
		version = "4.0.0"
		source  = "wrike.com/user/wrike"
	  }
	}
}
  
  provider "wrike"{
	  token =  "bearer ACCESS_TOKEN"
  }

resource "wrike_user" "user"{
    email = "xyz@gmail.com"
    role = "Collaborator"
    external = true
}

output "resourse_user"{
	  value = wrike_user.user
}

data "wrike_user" "user" {
    email = "xyz@gmail.com"
}

output "datasouce_user"{
    value = data.wrike_user.user
}
```

## Argument Reference

* `email` (string) - The email id associated with the user account.
* `firstname` (string) - First name of the User.
* `lastname` (string) - Last Name / Family Name / Surname of the User.
* `userid` (string) - unique id of a user.
* `accountid` (string) - account id of owner/admin account.
* `role` (string) - role of a user. Allowed values - User, Collaborator. 
* `external` (bool) - true for external user and Collaborator, false for Regular user.


## Exceptions

* No API for deleting the user and activating/deactivating the user.
* Not allowed to assign administrator role through API.
* But a administrator account can be changed to other roles through API.
* Not allowed to update Owner