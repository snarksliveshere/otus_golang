# 2019-11 there are events
# 2001-05 there are no events
Feature: Get Events for Month
  When I request GET events in Month "events-for-month" with param "month"
  I want to see the list of events and status successfully
  Or status error and text error

  Scenario: There are events for this month
    When I send "GET" request to router events-for-month "events-for-month" with param "month" and value "2019-11"
    Then The response code should be 200
    And The response should have length more than 0
    And status should be equal to success "success"

  Scenario: There are no events for this month
    When I send "GET" request to router events-for-month "events-for-month" there are no events with param "month" and value "2001-05"
    Then The response code should be 200
    And status should be equal to error "error"
    And The error text must be non empty string