package configs

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var Database *gorm.DB

func DatabaseDB() {
	database_mysql, err := gorm.Open("mysql", "test:test@tcp(127.0.0.1:3307)/MYSQL_REST_API_TEST")

	if err != nil {
		fmt.Println("NOT CONNECTED")
	}

	Database = database_mysql
}
