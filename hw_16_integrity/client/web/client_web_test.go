package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
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

func (test *notifyTest) returnPostValueResponse(httpMethod, addr, params string) (resp *http.Response, err error) {
	switch httpMethod {
	case http.MethodPost:
		resp, err = http.Post(addr, "application/x-www-form-urlencoded", strings.NewReader(params))
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (test *notifyTest) returnPostResponse(httpMethod, addr, contentType string, data *gherkin.DocString) (resp *http.Response, err error) {
	switch httpMethod {
	case http.MethodPost:
		replacer := strings.NewReplacer("\n", "", "\t", "", " ", "")
		cleanJson := replacer.Replace(data.Content)
		fmt.Printf("addr: %s, content-type: %s, params: %s\n", addr, contentType, cleanJson)
		resp, err = http.Post(addr, contentType, bytes.NewReader([]byte(cleanJson)))
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}
	i, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(i))
	fmt.Println("return post resp")
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

func (test *notifyTest) createUrlWithPostParams(router string, params map[string]string) (string, string, error) {
	apiStr := "http://" + conf.ListenIP + ":" + conf.WEBPort + "/" + router
	apiUrl, err := url.Parse(apiStr)
	if err != nil {
		return "", "", err
	}
	parameters := url.Values{}
	for k, v := range params {
		parameters.Set(k, v)
	}
	fmt.Printf("make url: %s, parameters: %s\n", apiUrl.String(), parameters.Encode())
	return apiUrl.String(), parameters.Encode(), nil
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

func (test *notifyTest) iSendRequestToRouterEventsfordayWithParamAndValue(httpMethod, router, param, value string) error {
	addr, err := test.createUrlWithGetParams(router, map[string]string{param: value})
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

func (test *notifyTest) unmarshallResponseToStruct(bytes []byte) error {
	resp := new(Response)
	err := json.Unmarshal(bytes, resp)
	if err != nil {
		fmt.Println(string(bytes))
		return err
	}
	test.responseStruct = resp
	return nil
}

func (test *notifyTest) theResponseShouldHaveLengthMoreThan(zeroLen int) error {
	err := test.unmarshallResponseToStruct(test.responseBody)
	if err != nil {
		return err
	}
	if len(test.responseStruct.Events) == zeroLen {
		return fmt.Errorf("unexpected events length: %d must be more than zero %d", len(test.responseStruct.Events), zeroLen)
	}
	return nil
}

func (test *notifyTest) statusShouldBeEqualToSuccess(status string) error {
	if test.responseStruct.Status != status {
		return fmt.Errorf("status must be: %s, not %s", status, test.responseStruct.Status)
	}
	return nil
}

func (test *notifyTest) iSendRequestToRouterEventsfordayThereAreNoEventsWithParamAndValue(httpMethod, router, param, value string) error {
	addr, err := test.createUrlWithGetParams(router, map[string]string{param: value})
	if err != nil {
		return err
	}
	resp, err := test.returnGetResponse(httpMethod, addr)
	if err != nil {
		return err
	}
	err = test.enrichTestStruct(resp)
	if err != nil {
		return err
	}
	return nil
}

func (test *notifyTest) enrichTestStruct(resp *http.Response) error {
	test.responseStatusCode = resp.StatusCode
	test.status = resp.Status

	rb, err := ioutil.ReadAll(resp.Body)
	err = test.unmarshallResponseToStruct(rb)
	if err != nil {
		return err
	}
	return nil
}

func (test *notifyTest) statusShouldBeEqualToError(status string) error {
	if test.responseStruct.Status != status {
		return fmt.Errorf("status must be: %s, not %s", status, test.responseStruct.Status)
	}
	return nil
}

func (test *notifyTest) theErrorTextMustBeNonEmptyString() error {
	if test.responseStruct.Error == "" {
		return fmt.Errorf("error must not be empty string: %s", test.responseStruct.Error)
	}
	return nil
}

func (test *notifyTest) iSendGoodRequestToRouterWithData(httpMethod, router, contentType string, data *gherkin.DocString) error {
	addr, _, err := test.createUrlWithPostParams(router, nil)
	if err != nil {
		return err
	}
	resp, err := test.returnPostResponse(httpMethod, addr, contentType, data)
	if err != nil {
		return err
	}

	err = test.enrichTestStruct(resp)
	if err != nil {
		return err
	}
	return nil
}

func (test *notifyTest) eventShouldExist() error {
	return godog.ErrPending
}

func (test *notifyTest) iSendBadRequestToRouterWithData(httpMethod, router, contentType string, data *gherkin.DocString) error {
	addr, _, err := test.createUrlWithPostParams(router, nil)
	if err != nil {
		return err
	}
	resp, err := test.returnPostResponse(httpMethod, addr, contentType, data)
	if err != nil {
		return err
	}
	err = test.enrichTestStruct(resp)
	if err != nil {
		return err
	}
	return nil
}

func (test *notifyTest) iSendGoodRequestToRouterWithDateTitleDescription(httpMethod, router, date, title, description string) error {
	mm := make(map[string]string, 3)
	mm["date"] = date
	mm["title"] = title
	mm["description"] = description
	url, params, err := test.createUrlWithPostParams(router, mm)
	if err != nil {
		return err
	}
	resp, err := test.returnPostValueResponse(httpMethod, url, params)
	if err != nil {
		return err
	}
	err = test.enrichTestStruct(resp)
	if err != nil {
		return err
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	test := new(notifyTest)
	// healthcheck
	s.Step(`^I send "([^"]*)" request to healthCheck "([^"]*)"$`, test.iSendRequestToHealthCheck)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^The response should match text "([^"]*)"$`, test.theResponseShouldMatchText)

	// create-event
	//s.Step(`^I send "([^"]*)" good request to router "([^"]*)" with "([^"]*)" data:$`, test.iSendGoodRequestToRouterWithData)
	s.Step(`^I send "([^"]*)" good request to router "([^"]*)" with date "([^"]*)" title "([^"]*)" description "([^"]*)"$`, test.iSendGoodRequestToRouterWithDateTitleDescription)
	s.Step(`^status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^event should exist$`, test.eventShouldExist)

	s.Step(`^I send "([^"]*)" bad request to router "([^"]*)" with "([^"]*)" data:$`, test.iSendBadRequestToRouterWithData)
	s.Step(`^status should be equal to error "([^"]*)"$`, test.statusShouldBeEqualToError)
	s.Step(`^The error text must be non empty string$`, test.theErrorTextMustBeNonEmptyString)

	// get events-for-day
	s.Step(`^I send "([^"]*)" request to router events-for-day "([^"]*)" with param "([^"]*)" and value "([^"]*)"$`, test.iSendRequestToRouterEventsfordayWithParamAndValue)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^The response should have length more than (\d+)$`, test.theResponseShouldHaveLengthMoreThan)
	s.Step(`^status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)

	s.Step(`^I send "([^"]*)" request to router events-for-day "([^"]*)" there are no events with param "([^"]*)" and value "([^"]*)"$`, test.iSendRequestToRouterEventsfordayThereAreNoEventsWithParamAndValue)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^status should be equal to error "([^"]*)"$`, test.statusShouldBeEqualToError)
	s.Step(`^The error text must be non empty string$`, test.theErrorTextMustBeNonEmptyString)

}
