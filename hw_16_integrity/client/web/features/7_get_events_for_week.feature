# 2019-11-07 - 2019-11-13 there are events
# 2001-05-02 - 2001-05-08 there are no events
Feature: Get Events for Week
  When I request GET events in Month "events-for-week" with param "from" & "till"
  I want to see the list of events and status successfully
  Or status error and text error

  Scenario: There are events for Week
    When I send "GET" request to router events-for-week "events-for-week" with "from" from "2019-11-06" and "till" till "2019-11-13"
    Then The response code should be 200
    And The response should have length more than 0
    And status should be equal to success "success"

  Scenario: There are no events for this month
    When I send "GET" request to router events-for-week "events-for-week" with "from" from "2001-05-02" and "till" till "2001-05-08"
    Then The response code should be 200
    And status should be equal to error "error"
    And The error text must be non empty string