package main

import (
	"restApiTest/route"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	route.URL()
}
