Feature: Get Events for one Day
    When I request events in one Day
    I want to see the list of events and status successfully
    Or status error and text error

    Scenario: There are events for this day
        When I send "GET" request to router events-for-day "events-for-day" with param "date" and value "2019-11-10"
        Then The response code should be 200
        And The response should have length more than 0
        And status should be equal to success "success"

    Scenario: There are not events for this day
        When I send "GET" request to router events-for-day "events-for-day" there are no events with param "date" and value "2019-05-01"
        Then The response code should be 200
        And status should be equal to error "error"
        And The error text must be non empty string