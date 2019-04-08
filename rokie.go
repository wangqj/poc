package main

import (
	_ "github.com/apache/servicecomb-rokie/handler"

	"github.com/apache/servicecomb-rokie/cmd"
	"github.com/apache/servicecomb-rokie/config"
	"github.com/apache/servicecomb-rokie/pkg/resource"
	"github.com/go-chassis/go-chassis"
	"github.com/go-mesh/openlogging"
)

func main() {
	if err := cmd.Init(); err != nil {
		openlogging.Fatal(err.Error())
	}
	chassis.RegisterSchema("rest", &resource.KVResource{})
	if err := chassis.Init(); err != nil {
		openlogging.Fatal(err.Error())
	}
	if err := config.Init(cmd.Configs.ConfigFile); err != nil {
		openlogging.Fatal(err.Error())
	}
	if err := chassis.Run(); err != nil {
		openlogging.Error("service exit: " + err.Error())
	}
}
