# Wrike Terraform Provider

* This terraform provider allows to perform Create, Read and Import. 
* To fetch and import a user, server generated a User ID is required.
* Wrike provide API for update the user but there is one parameter to request i.e. role. Only you can update the role of a user. But you can update and delete the user through console.
* Wrike doesn't provides an API to delete, activate/deactivate the user.

## Requirements

* [Go](https://golang.org/doc/install) 1.16 <br>
* [Terraform](https://www.terraform.io/downloads.html) 0.13.x <br/>
* [Wrike API Access](https://developers.wrike.com/api/v4/users/)
* Wrike Admin Account (Wrike Admin of organization can grant the Admin access to the user and can give API access credentials and also the Authentication permanent token through client id & client secret)

## Initialise Wrike Provider in local machine 
1. Run the following command :
 ```golang
   go mod init terraform-provider-wrike
   go mod vendor
 ```

2. Run `go mod vendor` to create a vendor directory that contains all the provider's dependencies. <br>

## Installation
1. Run the following command to create a vendor subdirectory
2. which will comprise of all provider dependencies. <br>
```
    ~/.terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${target}
``` 
Command: 
```bash
mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/wrike/4.0.0/[OS_ARCH]
```
For eg. `mkdir -p ~/terraform.d\plugins\wrike.com\user\wrike\4.0.0\windows_amd64`<br>

1. Run `go build -o terraform-provider-wrike.exe`. This will save the binary (`.exe`) file in the main/root directory. <br>
2. Run this command to move this binary file to the appropriate location.
 ```
 move terraform-provider-wrike.exe %APPDATA%\terraform.d\plugins\hashicorp.com\edu\wrike\4.0.0\[OS_ARCH]
 ``` 
Otherwise you can manually move the file from current directory to destination directory.<br>


[OR]

1. Download required binaries <br>
2. move binary `~/.terraform.d/plugins/[architecture name]/`


## Run the Terraform provider

### Authentication

1. Create an app from admin account.
2. Through client id and secret, create a `Permanent Access Token`.
3. Then, paste the token into the tf file.

#### Create User
1. Add the user email in the respective field in `main.tf`. We can invite the user through email only.
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that a user has been successfully created and an user seat claim mail has been sent to the user.

#### Read the User Data
Add data and output blocks in the `main.tf` file after that add email field  and user email and run `terraform plan` to read user data


#### Import a User Data
1. Write manually a resource configuration block for the User in `main.tf`, to which the imported object will be mapped or define the empty resource block.
2. Run the command `terraform import wrike_user.user “[EMAIL_ID]”`
3. Check for the attributes in the `.tfstate` file and fill them accordingly in the resource block.


### Testing the Provider
1. Navigate to the test file directory.
2. Run command `go test` for unit testing and for acceptance testing run command `go test` . These commands will give combined test results for the execution or errors if any failure occurs.
3. If you want to see test result of each test function individually while running test in a single go, run command `go test -v`
4. To check test cover run `go test -cover`


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

resource "wrike_resource_user" "user"{
  userid      = "[USER_ID]"
  email       = "[USER_EMAIL]"
  firstname   = "[USER_FIRST_NAME]"
  lastname    = "[USER_LAST_NAME]"
  accountid   = "[ACCOUNT_ID]"
}

 output "resourse_list"{
	  value=wrike_user.user
  }
```

## Argument Reference

* `email` (string) - The email id associated with the user account.
* `firstname` (string) - First name of the User.
* `lastname` (string) - Last Name / Family Name / Surname of the User.
* `userid` (string) - unique id of a user, where you can read the user detail.
* `accountid` (string) - account id of owner/admin account. 
