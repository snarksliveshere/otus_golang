# do not create 2019-05-01 - this date used in another test as not existed
Feature: Create Event for Date
    When I send POST request "create-event"
    With params Title, Description, Date
    I want to get back the event and status success
    Or status error and text error

    Scenario: Successfully created event for date
        When I send "POST" good request to router "create-event" with date "2019-04-10T20:03+0300" title "test_title" description "description"
        Then The response code should be 200
        And status should be equal to success "success"
        And event should exist