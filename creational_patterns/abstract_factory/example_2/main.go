package main

import "log"

func main() {
	app, _ := new(Application).SetOsFactory(WINDOWS)
	log.Println(app.CreateUI())
}
