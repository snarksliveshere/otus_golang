package web

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/client/config"
	"log"
	"testing"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var (
	conf config.AppConfig
	//cc   *grpc.ClientConn
)

func init() {
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")
	//var err error
	//cc, err = grpc.Dial(conf.ListenIP+":"+conf.GRPCPort, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("cant connect to host: %s, port: %s, error:%s", conf.ListenIP, conf.GRPCPort, err.Error())
	//}
}

func TestExample(t *testing.T) {
	cases := []struct {
		status, title, description, time string
	}{
		{
			status:      "success",
			title:       "some new title",
			description: "some new description",
			time:        "2019-12-07T20:03+0300",
		},
		{
			status:      "success",
			title:       "new title2",
			description: "some new description2",
			time:        "2019-12-07T19:03+0300",
		},
		{
			status:      "success",
			title:       "new title3",
			description: "some new description3",
			time:        "2019-12-07T19:30+0300",
		},
		{
			status:      "success",
			title:       "new title3",
			description: "some new description3",
			time:        "2019-12-08T11:33+0300",
		},
	}
	log.Printf("app env: %#v", conf)
	for _, c := range cases {
		if c.title != c.title {
			t.Errorf("TestExample() title: %s", c.title)
		}
	}
}
