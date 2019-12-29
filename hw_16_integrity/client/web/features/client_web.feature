# file: features/client_web.feature

# http://localhost:8888/

# http://registration_service:8888/

Feature: API events CRUD
	As API client of events service
	In order to understand that the user can get and create events
	I want to receive event from API requests

	Scenario: API is available
		When I send "GET" request to healthCheck "healthcheck"
		Then The response code should be 200
		And The response should match text "OK"