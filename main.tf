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

    email= "abhisheksingh17@ece.iiitp.ac.in"
  }

  data "wrike_user" "user"{
	  email="abhishek.s@clevertap.com"
  }

  output "list"{
	  value=data.wrike_user.user
  }

  output "resourse_list"{
	  value=wrike_user.user
  }
  
  