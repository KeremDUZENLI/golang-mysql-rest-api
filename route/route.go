package route

import (
	"net/http"
	"restApiTest/configs"
	"restApiTest/controller"
	"restApiTest/model"

	"github.com/gin-gonic/gin"
)

func URL() {
	// database
	configs.DatabaseDB()
	configs.Database.AutoMigrate(&model.User{})

	// router
	ginRouter := gin.Default()

	// controllerDemo.go
	ginRouter.GET("/begin", controller.Begin)
	ginRouter.POST("/create", controller.Create)
	ginRouter.GET("/read", controller.Read)
	ginRouter.PUT("/update/:employeeID", controller.Update)
	ginRouter.DELETE("/delete/:employeeID", controller.Delete)

	// controller.go
	ginRouter.GET("/beginDatabase", controller.BeginDatabase)
	ginRouter.POST("/createDatabase", controller.CreateDatabase)
	ginRouter.GET("/readDatabase", controller.ReadDatabase)
	ginRouter.PUT("/updateDatabase/:employeeID", controller.UpdateDatabase)
	ginRouter.DELETE("/deleteDatabase/:employeeID", controller.DeleteDatabase)

	// adress
	http.ListenAndServe(":8888", ginRouter)
}
