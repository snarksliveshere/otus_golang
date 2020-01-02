# do not create 2019-05-01 - this date used in another test as not existed
Feature: Update Event by eventId
    When I send POST request "update-event"
    With params Title, Description & eventId
    I want to get back the event and status success
    Or status error and text error

    Scenario: Successfully updated event by Id
        When I send "POST" update request to router "update-event" with title "updated_title" description "updated_description"
        Then The response code should be 200
        And status should be equal to success "success"
        And event should exist