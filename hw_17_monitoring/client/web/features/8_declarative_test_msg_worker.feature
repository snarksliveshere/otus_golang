Feature: Verify that the message sender is working.
  The verifying is declarative,
  Since it is not yet known how the monitoring of message sending occurs
  And where they are sent at all

  Scenario: When messages are sent, the msg status and message body is written to the database
    When I establish a connection to the database
    Then The rows should have length more than 0
