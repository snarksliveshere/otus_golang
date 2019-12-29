package web

import (
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/client/config"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

type Event struct {
	Id          uint64    `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Time        time.Time `json:"time"`
	DateFk      uint32    `json:"dateFk"`
}

type Date struct {
	Id            uint32 `json:"id"`
	Day           string `json:"day,omitempty"`
	Description   string
	IsCelebration bool
	Events        []Event `json:"events,omitempty"`
}

type Response struct {
	Date   Date    `json:"day,omitempty"`
	Event  Event   `json:"event,omitempty"`
	Events []Event `json:"events,omitempty"`
	Error  string  `json:"error,omitempty"`
	Status string  `json:"status,omitempty"`
}

type notifyTest struct {
	responseStatusCode int
	responseBody       []byte
	status, errorText  string
	events             []interface{}
	responseStruct     *Response
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
	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:        "pretty", // Замените на "pretty" для лучшего вывода
		Paths:         []string{"features"},
		Randomize:     0, // Последовательный порядок исполнения
		StopOnFailure: true,
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func (test *notifyTest) returnGetResponse(httpMethod, addr string) (resp *http.Response, err error) {
	switch httpMethod {
	case http.MethodGet:
		resp, err = http.Get(strings.TrimSpace(addr))
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (test *notifyTest) iSendRequestToHealthCheck(httpMethod, healthCheckRouter string) error {
	addr, err := test.createUrlWithGetParams(healthCheckRouter, nil)
	if err != nil {
		return err
	}
	//addr := "http://" + conf.ListenIP + ":" + conf.WEBPort + "/" + healthCheckRouter
	resp, err := test.returnGetResponse(httpMethod, addr)
	if err != nil {
		return err
	}
	test.responseStatusCode = resp.StatusCode
	test.responseBody, err = ioutil.ReadAll(resp.Body)
	return nil
}

func (test *notifyTest) createUrlWithGetParams(router string, params map[string]string) (string, error) {
	apiStr := "http://" + conf.ListenIP + ":" + conf.WEBPort + "/" + router
	apiUrl, err := url.Parse(apiStr)
	if err != nil {
		return "", err
	}
	if params == nil {
		fmt.Printf("make url: %s", apiUrl.String())
		return apiUrl.String(), nil
	}
	parameters := url.Values{}
	for k, v := range params {
		parameters.Add(k, v)
	}
	apiUrl.RawQuery = parameters.Encode()
	fmt.Printf("make url: %s", apiUrl.String())
	return apiUrl.String(), nil
}

func (test *notifyTest) theResponseCodeShouldBe(code int) error {
	if test.responseStatusCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.responseStatusCode, code)
	}
	return nil
}

func (test *notifyTest) theResponseShouldMatchText(responseText string) error {
	if string(test.responseBody) != responseText {
		return fmt.Errorf("unexpected text: %s != %s", test.responseBody, responseText)
	}
	return nil
}

func (test *notifyTest) iSendRequestToRouterEventsfordayWithDay(httpMethod, router, dayParam string) error {
	addr, err := test.createUrlWithGetParams(router, map[string]string{"date": "2019-11-10"})
	if err != nil {
		return err
	}
	resp, err := test.returnGetResponse(httpMethod, addr)
	if err != nil {
		return err
	}
	test.responseStatusCode = resp.StatusCode
	test.status = resp.Status

	test.responseBody, err = ioutil.ReadAll(resp.Body)
	return nil
}

func (test *notifyTest) theResponseShouldHaveLengthMoreThan(zeroLen int) error {
	resp := new(Response)
	err := json.Unmarshal(test.responseBody, resp)
	if err != nil {
		return err
	}
	if len(resp.Events) == zeroLen {
		return fmt.Errorf("unexpected events length: %d must be more than zero %d", len(resp.Events), zeroLen)
	}
	test.responseStruct = resp
	return nil
}

func (test *notifyTest) statusShouldBeEqualToSuccess(status string) error {
	fmt.Println("olala restarting..........................")
	if test.responseStruct.Status != status {
		return fmt.Errorf("status must be: %s, not %s", status, test.responseStruct.Status)
	}
	return nil
}

func (test *notifyTest) iSendRequestToRouterEventsfordayWithDayThereAreNoEvents(arg1, arg2, arg3 string) error {
	return godog.ErrPending
}

func (test *notifyTest) statusShouldBeEqualToError(arg1 string) error {
	return godog.ErrPending
}

func (test *notifyTest) theErrorTextMustBeNonEmptyString() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	test := new(notifyTest)
	s.Step(`^I send "([^"]*)" request to healthCheck "([^"]*)"$`, test.iSendRequestToHealthCheck)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^The response should match text "([^"]*)"$`, test.theResponseShouldMatchText)

	s.Step(`^I send "([^"]*)" request to router events-for-day "([^"]*)" with day "([^"]*)"$`, test.iSendRequestToRouterEventsfordayWithDay)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^The response should have length more than (\d+)$`, test.theResponseShouldHaveLengthMoreThan)
	s.Step(`^status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)

	s.Step(`^I send "([^"]*)" request to router events-for-day "([^"]*)" with day "([^"]*)" there are no events$`, test.iSendRequestToRouterEventsfordayWithDayThereAreNoEvents)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^status should be equal to error "([^"]*)"$`, test.statusShouldBeEqualToError)
	s.Step(`^The error text must be non empty string$`, test.theErrorTextMustBeNonEmptyString)

}
