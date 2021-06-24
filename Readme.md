This Terraform provider enables create, read operations for wrike users.

## Requirements

* [Go](https://golang.org/doc/install) >= 1.16 (To build the provider plugin)<br>
* [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x <br/>
* Application: [Wrike](https://www.wrike.com/) (API support some operation only i.e. read and create).


## Application Account

### Setup
1. Create an wrike account at https://www.wrike.com/<br>
2. Sign in to the wrike account and start a free trial.<br>
3. After creating an account, go to account management and can manage the users from console manually.<br>

### API Authentication
1. To authenticate API we need a pair of credentials: Client ID and Client secret.<br>
2. To access client id and client secret, you need create an `wrike App` from Apps and Integration<br>
3. A pair of credentials: Client ID and Client secret will be generated and also write one Redirect URI (eg https://www.wrike.com/).<br>
4. With the help of pair of credential, you can generate `Permanent Access Token`.<br>


## Building The Provider
1. Clone the repository and add the dependencies. For this run the following commands: <br>
```
git clone https://github.com/abhisheksingh3897/terraform-provider-wrike.git
cd terraform-provider-wrike
go mod init terraform-provider-wrike
go mod tidy
```
2. Run `go mod vendor` to create a vendor directory that contains all the provider's dependencies. <br>


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
 <p align="center">
 [OR]
 </p>
 
  Manually move the file from current directory to destination directory (`%APPDATA%\terraform.d\plugins\wrike.com\user\wrike\4.0.0\windows_amd64`).<br>


## Working with terraform

### Application Credential Integration in terraform
1. Add terraform block and provider block as shown in Example Usage below.
2. Get a pair of credentials: Client ID and Client secret. By help of this, generate access token.
3. Assign the above Access Token to the repective field in the provider block.

### Basic Terraform Commands
1. `terraform init` - To initialize a working directory containing Terraform configuration files.
2. `terraform plan` - To create an execution plan. Displays the changes to be done.
3. `terraform apply` - To execute the actions proposed in a Terraform plan. Apply the chages.

#### Create User
1. Add the user email in the respective field in `main.tf`. We can invite the user through email only.
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`.
4. You will see that a user has been successfully created and an user seat claim mail has been sent to the user.

#### Read the User Data
Add data and output blocks in the `main.tf` file after that add email field and user email and run `terraform plan` to read user data.

#### Import a User Data
1. Write manually a resource configuration block for the User in `main.tf`, to which the imported object will be mapped or define the empty resource block.
2. Run the command `terraform import wrike_user.user “[EMAIL_ID]”`
3. Run `terraform plan`, if output show `0 to addd, 0 to change and 0 to destroy` user import is successful, otherwise recheck the employee data in resource block with employee data in the policy in Expensify Website. 
4. Check for the attributes in the `.tfstate` file and fill them accordingly in the resource block.

## Example Usage

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
	  token =  "bearer eyJ0dCI6InAiLCJhbGciOiJIUzI1NiIsInR2IjoiMSJ9.eyJkIjoie1wiYVwiOjQ2NTAxODYsXCJpXCI6Nzc0MzU1NyxcImNcIjo0NjI2MjkwLFwidVwiOjEwNjI2NjMwLFwiclwiOlwiVVNcIixcInNcIjpbXCJXXCIsXCJGXCIsXCJJXCIsXCJVXCIsXCJLXCIsXCJDXCIsXCJEXCIsXCJNXCIsXCJBXCIsXCJMXCIsXCJQXCJdLFwielwiOltdLFwidFwiOjB9IiwiaWF0IjoxNjIxMTc5NTkyfQ.56vbcUlIBctouj49OcOQoID0ehSmq4DveZHjKX3J2jY"
  }

resource "wrike_user" "user"{
    email= "abhishek.singh3897@gmail.com"
  }

resource "wrike_user" "user"{
  userid      = "[USER_ID]"
  email       = "[USER_EMAIL]"
  firstname   = "[USER_FIRST_NAME]"
  lastname    = "[USER_LAST_NAME]"
  accountid   = "[ACCOUNT_ID]"
}

 output "resourse_list"{
	  value=wrike_user.user
  }

data "wrike_user" "user" {
    userid = "[USER_ID]"
}

output "datasouce_user"{
    value = data.wrike_user.user
}
```

## Argument Reference

* `email` (string) - The email id associated with the user account.
* `firstname` (string) - First name of the User.
* `lastname` (string) - Last Name / Family Name / Surname of the User.
* `userid` (string) - unique id of a user, where you can read the user detail.
* `accountid` (string) - account id of owner/admin account. 


## Exceptions

* Update API should not update anything of the user. Wrike update their user details through console.
* Also no API for deleting the user and activating/deactivating a user.
* To read the user data, we need user USER_ID for fetch the user detail. You can get the USER_ID from the contact API. `https://www.wrike.com/api/v4/contacts`.