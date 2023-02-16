package controller

import (
	"net/http"
	"restApiTest/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BeginDatabase(c *gin.Context) {
	userBegin := &model.User{}
	c.ShouldBindJSON(userBegin)

	c.IndentedJSON(http.StatusOK, userBegin.DatabaseBeginUser())
}

func CreateDatabase(c *gin.Context) {
	userCreate := &model.User{}
	c.ShouldBindJSON(userCreate)

	c.IndentedJSON(http.StatusOK, userCreate.DatabaseCreateUser())
}

func ReadDatabase(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.DatabaseReadUser())
}

func UpdateDatabase(c *gin.Context) {
	nStr := c.Param("employeeID")
	n, _ := strconv.Atoi(nStr)

	userUpdate := &model.User{}
	c.ShouldBindJSON(userUpdate)

	for _, user := range *model.DatabaseReadUser() {
		if user.EmployeeID == n {
			model.DatabaseDeleteUser(n)

			c.IndentedJSON(http.StatusOK, user)
			c.IndentedJSON(http.StatusOK, userUpdate.DatabaseUpdateUser())
			break
		}
	}

}

func DeleteDatabase(c *gin.Context) {
	nStr := c.Param("employeeID")
	n, _ := strconv.Atoi(nStr)

	for _, user := range *model.DatabaseReadUser() {
		if user.EmployeeID == n {
			c.IndentedJSON(http.StatusOK, user)
			model.DatabaseDeleteUser(n)
			break
		}
	}
}
