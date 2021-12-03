package main

import (
	"workuo/driver"
	"workuo/routes"
)

func main() {
	driver.InitDB()
	e := routes.New()
	e.Start(":8000")
}
