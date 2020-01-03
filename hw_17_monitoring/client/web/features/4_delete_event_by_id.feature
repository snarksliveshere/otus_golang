# normally, i get eventId from outside, in this case, i get eventId from previous step
Feature: Delete Event by Id
    When I send to "delete-event" with method POST with param "eventId"
    I want to see status "success"
    Or status "error" and text error

    Scenario: There is id for this event
        When I send "POST" request to router delete-event "delete-event" with param "eventId"
        Then The response code should be 200
        And status should be equal to success "success"

    Scenario: There are no event for this id
        When I send "POST" request to router delete-event "delete-event" with not exist param "eventId"
        Then The response code should be 200
        And status should be equal to error "error"
        And The error text must be non empty string