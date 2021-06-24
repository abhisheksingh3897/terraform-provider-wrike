terraform {
	required_providers {
	  wrike = {
		version = "4.0.0"
		source  = "wrike.com/user/wrike"
	  }
	}
}
  
  provider "wrike" {
	  client_id = "ZtrqchYX"
    client_secret = "X9SLPR2tiZrsUmiqht1R0a8K6qpuYwrFaFBRJ2JikSJnIAPETv6cKFDvhJooBxU0"
    refresh_token = "eyJ0dCI6InAiLCJhbGciOiJIUzI1NiIsInR2IjoiMSJ9.eyJkIjoie1wiYVwiOjQ3NTIzODUsXCJpXCI6NzgwMzc0OCxcImNcIjo0NjI3MTA4LFwidlwiOlwiXCIsXCJ1XCI6MTA4ODk1MDgsXCJyXCI6XCJVU1wiLFwic1wiOltcIk5cIl0sXCJ6XCI6W1wicnNoXCJdLFwidFwiOjE2MjQ1MDcwMzUwMDB9IiwiZXhwIjoxNjI0NTA3MDM1LCJpYXQiOjE2MjQ1MDM0MzV9.eJ7BZf1-hAHX9aQOJQh3gc7ccXI9C4WaSEZFClIuolI"
  }
  
  resource "wrike_user" "user"{
    # role="Collaborator"
    # external=true
    email= "abhishek.singh3897.com"
  }

  # data "wrike_user" "user"{
	#   email="abhishek.s@clevertap.com"
  # }

  # output "list"{
	#   value=data.wrike_user.user
  # }

  output "resourse_user"{
	  value=wrike_user.user
  }
  
  