package client

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("WRIKE_TOKEN", "bearer eyJ0dCI6InAiLCJhbGciOiJIUzI1NiIsInR2IjoiMSJ9.eyJkIjoie1wiYVwiOjQ2NTAxODYsXCJpXCI6Nzc0MzU1NyxcImNcIjo0NjI2MjkwLFwidVwiOjEwNjI2NjMwLFwiclwiOlwiVVNcIixcInNcIjpbXCJXXCIsXCJGXCIsXCJJXCIsXCJVXCIsXCJLXCIsXCJDXCIsXCJEXCIsXCJNXCIsXCJBXCIsXCJMXCIsXCJQXCJdLFwielwiOltdLFwidFwiOjB9IiwiaWF0IjoxNjIxMTc5NTkyfQ.56vbcUlIBctouj49OcOQoID0ehSmq4DveZHjKX3J2jY")
}

func TestClient_GetItem(t *testing.T) {
	testCases := []struct {
		testName     string
		itemName     string
		expectErr    bool
		expectedResp *User
	}{
		{
			testName:  "user exists",
			itemName:  "abhishek.singh3897@gmail.com",
			expectErr: false,
			expectedResp: &User{
				ID:        "KUAKIUXG",
				FirstName: "abhi",
				LastName:  "s",
				Profile: []UserProfile{
					{
						Email:     "abhishek.singh3897@gmail.com",
						AccountID: "IEAEN5GK",
						Role:      "User",
						External:  false,
					},
				},
			},
		},
		{
			testName:     "user does not exist",
			itemName:     "ashutosh.verma@clevertap.com",
			expectErr:    true,
			expectedResp: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("WRIKE_TOKEN"))
			item, err := client.GetUser(tc.itemName)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, item)
		})
	}
}

func TestClient_NewUser(t *testing.T) {
	testCases := []struct {
		testName  string
		newItem   string
		expectErr bool
	}{
		{
			testName:  "success",
			newItem:   "ashutosh.verma@clevertap.com",
			expectErr: false,
		},
		{
			testName:  "item already exists",
			newItem:   "abhishek.singh3897@gmail.com",
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("WRIKE_TOKEN"))
			err := client.NewUser(tc.newItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestClient_UpdateItem(t *testing.T) {
	testCases := []struct {
		testName     string
		itemName     string
		accountId    string
		role         string
		external     bool
		expectErr    bool
		expectedResp *User
	}{
		{
			testName:  "user exists",
			itemName:  "abhisheksingh17@ece.iiitp.ac.in",
			accountId: "IEAEN5GK",
			role:      "Collaborator",
			external:  true,
			expectErr: false,
			expectedResp: &User{
				ID:        "KUAKMCN5",
				FirstName: "abhi",
				LastName:  "s",
				Profile: []UserProfile{
					{
						Email:     "abhisheksingh17@ece.iiitp.ac.in",
						AccountID: "IEAEN5GK",
						Role:      "Collaborator",
						External:  true,
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("WRIKE_TOKEN"))
			err := client.UpdateUser(tc.itemName, tc.accountId, tc.role, tc.external)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
