package main

import (
	"workuo/driver"
	"workuo/routes"
)

func main() {
	driver.InitDB()
	e := routes.New()
	e.Start("127.0.0.1:8000")
}
