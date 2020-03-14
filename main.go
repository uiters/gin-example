package main

import (
	"github.com/sirupsen/logrus"
	"mgo-gin/app"
	"mgo-gin/db"
)

func main(){
	var server app.Routes
	resource, err := db.InitResource()
	if err != nil {
		logrus.Error(err)
	}
	defer resource.Close()
	server.StartGin()
}