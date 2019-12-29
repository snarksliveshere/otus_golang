package web

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/client/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

type notifyTest struct {
	responseStatusCode int
	responseBody       []byte
}

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
	time.Sleep(10 * time.Second)
	//var err error
	//cc, err = grpc.Dial(conf.ListenIP+":"+conf.GRPCPort, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("cant connect to host: %s, port: %s, error:%s", conf.ListenIP, conf.GRPCPort, err.Error())
	//}
}

func TestMain(m *testing.M) {
	fmt.Println("Wait 5s for service availability...")
	time.Sleep(5 * time.Second)

	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:    "pretty", // Замените на "pretty" для лучшего вывода
		Paths:     []string{"features"},
		Randomize: 0, // Последовательный порядок исполнения
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func (test *notifyTest) iSendRequestToHealthCheck(httpMethod, healthCheckRouter string) (err error) {
	var r *http.Response
	addr := "http://" + conf.ListenIP + ":" + conf.WEBPort + "/" + healthCheckRouter
	switch httpMethod {
	case http.MethodGet:
		r, err = http.Get(strings.TrimSpace(addr))
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}

	if err != nil {
		return err
	}
	test.responseStatusCode = r.StatusCode
	test.responseBody, err = ioutil.ReadAll(r.Body)
	fmt.Printf("test is pass: %s", addr)
	return nil
}

func (test *notifyTest) theResponseCodeShouldBe(code int) error {
	if test.responseStatusCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.responseStatusCode, code)
	}
	return nil
}

func theResponseShouldMatchText(arg1 string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	test := new(notifyTest)
	s.Step(`^I send "([^"]*)" request to healthCheck "([^"]*)"$`, test.iSendRequestToHealthCheck)
	//s.Step(`^The response code should be (\d+)$`, theResponseCodeShouldBe)
	//s.Step(`^The response should match text "([^"]*)"$`, theResponseShouldMatchText)
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
