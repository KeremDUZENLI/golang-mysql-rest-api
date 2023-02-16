package model

import (
	"restApiTest/configs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DatabaseBeginUser(t *testing.T) {
	aStruct := User{}
	result := aStruct.DatabaseBeginUser()

	// assert.NotNil(t, result, "Expected a non-nil result")
	assert.Equal(t, &Users, result, "Expected result to be equal to Users slice")
}

func Test_DatabaseCreate(t *testing.T) {
	user := User{EmployeeID: 4, Name: "D", Email: "DD@", Document: "DDD", RegisteredAt: ""}

	// Check if the record is not found in the database
	assert.True(t, configs.Database.Where("employee_id = ?", user.EmployeeID).First(&user).RecordNotFound(), "Expected record not found")

	// Create the user in the database
	configs.Database.Create(&user)

	// Check if the record is now found in the database
	assert.False(t, configs.Database.Where("employee_id = ?", user.EmployeeID).First(&user).RecordNotFound(), "Expected record found")
}
